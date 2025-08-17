GOFMT ?= gofumpt -l -s
GO ?= go
PACKAGES ?= $(shell $(GO) list ./...)
SOURCES ?= $(shell find . -name "*.go" -type f)

all: lint

fmt:
	@hash gofumpt > /dev/null 2>&1; if [ $$? -ne 0 ]; then \
		$(GO) install mvdan.cc/gofumpt; \
	fi
	$(GOFMT) -w $(SOURCES)

vet:
	$(GO) vet $(PACKAGES)

test:
	@$(GO) test -v -cover -coverprofile coverage.txt $(PACKAGES) && echo "\n==>\033[32m Ok\033[m\n" || exit 1

clean:
	go clean -x -i ./...
	rm -rf coverage.txt $(EXECUTABLE) $(DIST) vendor

ssh-server:
	adduser -h /home/test-user -s /bin/sh -D -S test-user
	echo test-user:1234 | chpasswd
	mkdir -p /home/test-user/.ssh
	chmod 700 /home/test-user/.ssh
	cat tests/.ssh/id_rsa.pub >> /home/test-user/.ssh/authorized_keys
	cat tests/.ssh/test.pub >> /home/test-user/.ssh/authorized_keys
	chmod 600 /home/test-user/.ssh/authorized_keys
	chown -R test-user /home/test-user/.ssh
	# add public key to root user
	mkdir -p /root/.ssh
	chmod 700 /root/.ssh
	cat tests/.ssh/id_rsa.pub >> /root/.ssh/authorized_keys
	cat tests/.ssh/test.pub >> /root/.ssh/authorized_keys
	chmod 600 /root/.ssh/authorized_keys
	# Append the following entry to run ALL command without a password for a user named test-user:
	mkdir -p /etc/sudoers.d
	cat tests/sudoers >> /etc/sudoers.d/sudoers
	# install ssh and start server
	apk add --update openssh openrc dos2unix sudo
	dos2unix tests/entrypoint.sh
	rm -rf /etc/ssh/ssh_host_rsa_key /etc/ssh/ssh_host_dsa_key
	sed -i 's/^#PubkeyAuthentication yes/PubkeyAuthentication yes/g' /etc/ssh/sshd_config
	sed -i 's/AllowTcpForwarding no/AllowTcpForwarding yes/g' /etc/ssh/sshd_config
	sh tests/entrypoint.sh /usr/sbin/sshd -D &
