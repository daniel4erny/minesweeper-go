# Maintainer: dandul <externals@yordstudio.com>
pkgname=minesweeper-go-git
pkgver=r5.c6df0a0
pkgrel=1
pkgdesc="Terminal-based Minesweeper game written in Go"
arch=('x86_64' 'aarch64')
url="https://github.com/daniel4erny/minesweeper-go"
license=('Unlicense')
depends=()
makedepends=('go' 'git')
provides=('minesweeper-go')
conflicts=('minesweeper-go')
source=("$pkgname::git+https://github.com/daniel4erny/minesweeper-go.git")
sha256sums=('SKIP')

pkgver() {
	cd "$pkgname"
	printf "r%s.%s" "$(git rev-list --count HEAD)" "$(git rev-parse --short HEAD)"
}

build() {
	cd "$pkgname"
	export CGO_ENABLED=0
	export GOFLAGS="-buildmode=pie -trimpath -mod=readonly -modcacherw"
	go build -o minesweeper-go .
}

package() {
	cd "$pkgname"
	install -Dm755 minesweeper-go "$pkgdir/usr/bin/minesweeper-go"
}
