/*
Copyright 2018 Google LLC

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package secrets

import (
	"testing"
)

const (
	// Intentionally hard-coded private keys for testing
	/* #nosec */
	privNoPassphrase = `-----BEGIN PGP PRIVATE KEY BLOCK-----
Version: BCPG C# v1.6.1.0

lQOsBFv+Y7UBCADGf1/XErc1VjjakMwI5kJumfS+FpzFhIq9MsdzoOOD+o+q7Noj
0r5Y4acp9AGvA0fA7H15JdyT4DEEcNzUNyQySV/Huto0NT5t1L8JLI70/RhF38LA
kdSR+Q7Uhf/7+6zTB0nfRnc2nfp24PvWSuUlP8NsgUA6WCFY70w6S2iNnv0WsiLU
XoHpdPm1ke3LABfc5Ujv/4V38WrGb17RP9y3J1TS7TG9tb4ndAIOSEIoxB2njVp1
B76ZZWD6WbDxAr4V8X5CBEwyDV4mUE3fcouZXeOgFw5WoLIC14Q3YY1zOoS1SOCY
0joJZTFfoAlWYorbG4XREcSsSodsvEOFjLInABEBAAH/AwMC2dqsriFMWglgurJl
xHcugu8aqD9i2lPLnSGZyRNV9Wu3Uhwjj0ROTjVm7mcTwxaCEtk9UfzfIZ67SFmn
EFze8ZmNaOvlRNkwAOa3TtJbGuHA27HQyzlcRZ2lqYl9ABkUEj1hDrkk7lDN7qml
QA3JP3UCsuFwwSyBOp2XbbLBYAj/wgvJ/K3HaCp0xHHnQofAgROV1ccZXjIgwknh
iEx3cW7ga4LSVsC0U9gulOeOu4DjSohyP+NC+fE658VTYJvQ0NJoMhGaewqLskye
wjbSD1ZCABCW5GrgsJv/cqrA8EFehGXJVNCZWiW0w1kXMV+Lk60ZA12WycHYyVSi
3tvZtkj0VZc36TtEYp7YLkzsYVXgtVJOn0GfwYRUqxcI5nVqP2sv/uQ3h9PMvpo7
DRLxn8d/e9fwT4Cbxf0wQcu4T07QP0Nwi4btfDb4HeaiiKK8uoSZjUg542Ts42bN
iGiQ2eer0bU/gcH6LvczlQMXTuUfeZvrxf9h1djwZwKmi926vIGYFE96J+w9xwsz
hgSrONnDoj1ciKiRANBLV6ddDWiTNm4FQZAR1FZIrHwrzb34cJPEQvGI3yJfWVIS
ZmkELKF6W4d1e3lMsKneZJCD0ZI8uhb7m3flj1bx4NzJshmPDrw3/Zw6J7xZnL78
Yf6FtjYvdHifueYpIgt1IqvnjDrlBT57FRWwUlV/EP6hORylq2byEOJDT9NGJR9r
EuaHQGQTdLn6yeDaUOrLq7ei+O+3qU9+5ztpgTpvjgUiOpgbAQ2K13882YHjnG7k
l2JlEsLKZLfov5pbhpW4wzWIByfzSjCcSCW5jfRpxu0tkheX3dCm5/BaPXgPnHtN
8t/5x3gWhojEeF4mBB0EiemIRIO3zMh6rfZuiqNMo7QMdGVzdEBwZ3AuY29tiQEc
BBABAgAGBQJb/mO1AAoJEDbFvgxBN1W2do0IAKFb2YcATgHg/h3HaV2NKnwY2+qf
BrvMifI964XOe+KUsUn8lzdKdO/HqgaAY8pwtnWi/ZIKKwEmFLUJ4nnw5KCOYANH
98br3JLVXolSJET3HVQjLO66dqBwnxfRBE+RodI/xgN3hhKpY80TmSCV4c+8ZS9K
uk1OYjmhy7o41Uythm3JZtlVWaQwNrxuW8FZaKf+ddAunVpYdAxyGIvQOeSvvjT8
wS/ESDp9/HoTDeQ/xj2yTr7taoZNnNDe1WrQueqE/kHpB3+a6jsdfMmduP7KCk1c
SfFhCkm3GY7opmHqvdkoBSpk7DwEmao+Q3frLO7EFr8EZ4o8PfFay4+QMFM=
=zZoX
-----END PGP PRIVATE KEY BLOCK-----`

	/* #nosec */
	privPassphrase = `-----BEGIN PGP PRIVATE KEY BLOCK-----

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
=1siz
-----END PGP PRIVATE KEY BLOCK-----`

	public = `-----BEGIN PGP PUBLIC KEY BLOCK-----
Version: BCPG C# v1.6.1.0

mQENBFv+Y7UBCADGf1/XErc1VjjakMwI5kJumfS+FpzFhIq9MsdzoOOD+o+q7Noj
0r5Y4acp9AGvA0fA7H15JdyT4DEEcNzUNyQySV/Huto0NT5t1L8JLI70/RhF38LA
kdSR+Q7Uhf/7+6zTB0nfRnc2nfp24PvWSuUlP8NsgUA6WCFY70w6S2iNnv0WsiLU
XoHpdPm1ke3LABfc5Ujv/4V38WrGb17RP9y3J1TS7TG9tb4ndAIOSEIoxB2njVp1
B76ZZWD6WbDxAr4V8X5CBEwyDV4mUE3fcouZXeOgFw5WoLIC14Q3YY1zOoS1SOCY
0joJZTFfoAlWYorbG4XREcSsSodsvEOFjLInABEBAAG0DHRlc3RAcGdwLmNvbYkB
HAQQAQIABgUCW/5jtQAKCRA2xb4MQTdVtnaNCAChW9mHAE4B4P4dx2ldjSp8GNvq
nwa7zInyPeuFznvilLFJ/Jc3SnTvx6oGgGPKcLZ1ov2SCisBJhS1CeJ58OSgjmAD
R/fG69yS1V6JUiRE9x1UIyzuunagcJ8X0QRPkaHSP8YDd4YSqWPNE5kgleHPvGUv
SrpNTmI5ocu6ONVMrYZtyWbZVVmkMDa8blvBWWin/nXQLp1aWHQMchiL0Dnkr740
/MEvxEg6ffx6Ew3kP8Y9sk6+7WqGTZzQ3tVq0LnqhP5B6Qd/muo7HXzJnbj+ygpN
XEnxYQpJtxmO6KZh6r3ZKAUqZOw8BJmqPkN36yzuxBa/BGeKPD3xWsuPkDBT
=4l+X
-----END PGP PUBLIC KEY BLOCK-----`
)

var secrets = []struct {
	name       string
	private    string
	passphrase string
	public     string
}{
	{"no-passphrase", privNoPassphrase, "", public},
	{"passphrase", privPassphrase, "test passphrase", public},
}

func TestPgpKey(t *testing.T) {
	for _, s := range secrets {
		t.Run(s.name, func(t *testing.T) {
			key, err := NewPgpKey(s.private, s.passphrase, s.public)
			if err != nil {
				t.Error(err)
			}
			if key == nil {
				t.Fatalf("Got nil key")
			}
			if key.publicKey == nil {
				t.Fatalf("Got nil public key")
			}
			if key.privateKey == nil {
				t.Fatalf("Got nil private key")
			}
		})
	}
}

func TestKeyAndFingerprint(t *testing.T) {
	// Public key below is obtained from the integration test key:
	// $ mkdir /tmp/h
	// $ gpg --homedir=/tmp/h --import integration/testdata/keys/attestor-1-secret-key.pgp
	// $ gpg --homedir=/tmp/h --armor --export attestor-1@example.com | base64 -w 0
	publicKey := `LS0tLS1CRUdJTiBQR1AgUFVCTElDIEtFWSBCTE9DSy0tLS0tCgptUUVOQkYzejNKc0JDQURTUStEU2p1Rkw2ME5mTGt6cUpyZDJESHAydHdqTkx5Y05zMGhBSkw1YUpsVU9JYWd4CkZ5QmRlb2VWSWxTSUV1LzN4UjA4V2RpUmJDZDNPbU8zbm1BTHB6Ti9hOUpDVzlCUXBTQ2Ftc3l6cHJZa2RzMzEKSGhTSzNHTC9PeVd6VzhKYzVMaFlNQzN2ZmJjQjJTc0E1WTdwU0dzcndtKy8wMHlKNU5BUzk2NERDbllBS0c5SgpWMk5QU01rZWU0em0vUEllUklaK2laVmFydXEzSDVCNGNIZnlQOHBzSE1hVnVMbVlPYnJPaTlPc1BzcnNlQk9RClExaVlEZnFpZytNVEdGbVNDZEhBM2RNTXFucHhKcW5sMHhtY0pjWHJrSXFrWXJsNVFaOElZMGxucUUyRnNjc24KZklFbDNZdlpDYXQxanI4SVE2L0hRZDhKcER4MUNHOGN4Nm5qQUJFQkFBRzBMQ0owWlhOMExXRjBkR1Z6ZEc5eQpMVEVpSUR3aVlYUjBaWE4wYjNJdE1VQmxlR0Z0Y0d4bExtTnZiU0kraVFGT0JCTUJDZ0E0RmlFRXNQUkNQRkNWCmZ0MUI2OFJycW9uem5kWkNtU0FGQWwzejNKc0NHeThGQ3drSUJ3SUdGUW9KQ0FzQ0JCWUNBd0VDSGdFQ0Y0QUEKQ2drUXFvbnpuZFpDbVNBRzJBZjlIWXViRUs3dW1ER2RHQzcvejV3L3pPREtERUZDNVZySU94MEJkU1RuM1JPSwpXdTR5S3hHZzR6NDVWTjZVL1d4NGYzbitOcTBab085MTBrZWVheGN6VmNTdVBibTEzMkZoM21VblJZZURXc0RtCmY4bFRwTWowMUVzQW5PZTJmWDdrOUZSRFl3TCtGZmR6VSt1RUhML243Nmd0MFBoKy9CbXdZbUUveWlUVjR2S2gKajhMS2pPSHJ0TUpDQm1GcU5UOGlhcnhUVXUrblBWMFc3SFlaTHFWOEpVdFVXRVdYbDlBVmxmZDJ1TkV6a1hqcApoMTloOVlMdEt3czFBM0lGZXQwS0dDcTVCdFo3WkdqdnMyQ0pqaGcvV3Z2YStNWXgyelVlRnVNRWlNZ1VGSTJZCkdjL3hMbGthTUtaQmRBcWdyaVBnL2FtVDBqVzlqdFFmcHNaZHQxQXhmUT09Cj1LMko3Ci0tLS0tRU5EIFBHUCBQVUJMSUMgS0VZIEJMT0NLLS0tLS0K`
        _, fingerprint, err := KeyAndFingerprint(publicKey)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if fingerprint != "B0F4423C50957EDD41EBC46BAA89F39DD6429920" {
		t.Errorf("Unexpected fingerprint: %v", fingerprint)
	}
}
