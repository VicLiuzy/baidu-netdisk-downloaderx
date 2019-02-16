package goproxy

import (
	"crypto/tls"
	"crypto/x509"
)

func init() {
	if goproxyCaErr != nil {
		panic("Error parsing builtin CA " + goproxyCaErr.Error())
	}
	var err error
	if GoproxyCa.Leaf, err = x509.ParseCertificate(GoproxyCa.Certificate[0]); err != nil {
		panic("Error parsing builtin CA " + err.Error())
	}
}

var tlsClientSkipVerify = &tls.Config{InsecureSkipVerify: true}

var defaultTLSConfig = &tls.Config{
	InsecureSkipVerify: true,
}

var CA_CERT = []byte(`-----BEGIN CERTIFICATE-----
MIID8zCCAtugAwIBAgIJALDpqK/0G0RrMA0GCSqGSIb3DQEBCwUAMIGPMQswCQYD
VQQGEwJDTjEPMA0GA1UECAwGWXVubmFuMRAwDgYDVQQHDAdLdW5taW5nMRowGAYD
VQQKDBFCM2xvZyBPcGVuIFNvdXJjZTERMA8GA1UECwwIQk5EIFRlYW0xEjAQBgNV
BAMMCWIzbG9nLm9yZzEaMBgGCSqGSIb3DQEJARYLZEBiM2xvZy5vcmcwHhcNMTgw
NDMwMTQ1OTUyWhcNMjgwNDI3MTQ1OTUyWjCBjzELMAkGA1UEBhMCQ04xDzANBgNV
BAgMBll1bm5hbjEQMA4GA1UEBwwHS3VubWluZzEaMBgGA1UECgwRQjNsb2cgT3Bl
biBTb3VyY2UxETAPBgNVBAsMCEJORCBUZWFtMRIwEAYDVQQDDAliM2xvZy5vcmcx
GjAYBgkqhkiG9w0BCQEWC2RAYjNsb2cub3JnMIIBIjANBgkqhkiG9w0BAQEFAAOC
AQ8AMIIBCgKCAQEAmWDYQ/RouBjky7f9skCinIiUDGsu7J4ef8ZPkJYtkWN4Nhjp
PUxOQwBoLuNmP3M5rqmJjZ3MkpGouLiKpp4psz37SMLzpmYIdgPFoiroOuTYAjW0
/qV0ggr7Jf5p9PDQyXg+/8Lz/7KsX0Wrl1whwE38xjHxvSvK9edMfS1KcU3emNaM
tmiMsfMyLy9Uq3hgNLEDF7a8WDXhs/WpkvmCLyuZYocgfI8ZLQGUYPtRMSzquvWL
wkDk7eKoTUiuf/i2EbrxjjJh1MEmUH+CfSkzWKSDJo64wDormzgasBWzeip9xZsU
QrZsE0Uu/6oFxaD75SfKnrcJ09bUgF6bWOcfPwIDAQABo1AwTjAdBgNVHQ4EFgQU
ypSfpPDl78WLySAv8TL35fbV2AQwHwYDVR0jBBgwFoAUypSfpPDl78WLySAv8TL3
5fbV2AQwDAYDVR0TBAUwAwEB/zANBgkqhkiG9w0BAQsFAAOCAQEAX+DfLMgFKzdm
zFDkB/Zgf5/tfllK2MdJlDstsL2TiPkUdzKBI5uFW1UfMm/oPCMSNzOLYtF0MzyY
EK8onuUXkixMiLLTNXZdREWQp23Kuprh3DtRP13oGp6Is2LQ/TM/k24j5gbUV7WC
B7KVcb+fClPUyhzT3LjsDlRElL+VZ65pONllTiuD0knmGH28kPuLUVFLGny3HpbH
ZI05Vjilos5g7pOsXhgh2bV33e1usLyenCLPw8QhLNIP73hu5Rp3L8CF7bZ9Zpk3
JXEGvNABa/IJRDy9/MQSPtkaxGwti2gpQ3cD/JAcY80woInCVg6siIVk5ErsWqRg
vExE2ipVog==
-----END CERTIFICATE-----`)

var CA_KEY = []byte(`-----BEGIN RSA PRIVATE KEY-----
MIIEowIBAAKCAQEAmWDYQ/RouBjky7f9skCinIiUDGsu7J4ef8ZPkJYtkWN4Nhjp
PUxOQwBoLuNmP3M5rqmJjZ3MkpGouLiKpp4psz37SMLzpmYIdgPFoiroOuTYAjW0
/qV0ggr7Jf5p9PDQyXg+/8Lz/7KsX0Wrl1whwE38xjHxvSvK9edMfS1KcU3emNaM
tmiMsfMyLy9Uq3hgNLEDF7a8WDXhs/WpkvmCLyuZYocgfI8ZLQGUYPtRMSzquvWL
wkDk7eKoTUiuf/i2EbrxjjJh1MEmUH+CfSkzWKSDJo64wDormzgasBWzeip9xZsU
QrZsE0Uu/6oFxaD75SfKnrcJ09bUgF6bWOcfPwIDAQABAoIBAEHvvklIcxRhr0pc
6LWOz9lXZb4I/f0+5IGtsVrJFQHhrzYic3KksShgrmkp3TT9xhhuTfIRmnaQI+Z1
fzWt3ONvi+110eUBDDYf3QCgTYG0C6C3kf2/B104d0uYGdjxBfD2vq6ZEugj2bKD
cwb41pGjdT2HDxNyFiisrNdLdytJ0jbuO+3h8RuOUxNnQ4ENRPj0YRoijk1Cuj3Y
0tbgPQC5R+OwTfXb8U0YFDwnWyr7xLhcX/64AqEsvgCx3+X3K59SWg51cJK7o27Q
eQEJHHsU8vu3jiBA3M24cvle4XshZBxTVS+ozlc3aRxGiG7zm6ppMMm5H4gdj6Sz
lHfkm2ECgYEAyeUvPcq8UoZp11DWFKNy67/zWqJRWwjogztStQOZwKFaJdw0vIkP
6KaeFBBPvCh4dUFoF008zmxDOMpfMVqp5+lEhvd62AfN9JoI77QPXLp5bPBJrn66
wn30jjVvP/JlMAnqlXWNhsDq6rHavpb4YxEUr1szZvX8qBsnXjf8ubkCgYEAwns0
biiWz4lndPw9YgxERyrbMDRgfNfcgJzYYIb8kfNiHoWTMqD0+ZjuOF3wqb91yJ03
PXAp7fhjPJQDhmpRf6WsW9yF73UjHUOl/OH6ef1KMq7Evdv1FoOCv1QEJHXhenV+
NzW1NTWDwJcohKW1RqaiiIvqVl25ZFQtQFTvPLcCgYAenLqCAwh09t57gE1Fidb1
+f/bITt2sy4WH9uViLs6Yvy5Ot9gKuQXoZPcdPCyFezV0poG1nY+0lmK9dxjLl0O
xebXFBUzRlOrH+P0/HdL1xZsNy7k32BP9ysqmN2Kc6V2GDrSeKWB5iR4e4kwVtHI
q6FYSVrM5MVvc9aPSavvOQKBgAMkgIVjLISwALb4OHm0HwpTW/Vcc06xIkQFxM1S
gRROrYA2wlkoFq1N1lOc2P21NU8Hk1Pv0w2gXOImtiPe5fA9Ghrl5lgnOWkIpc7/
S8/a0u8cudsjicNSXrN+xQ5dxUX4ThDQUlxIwBdFmgcCsnowySGRuTiqDQvc/Fo0
xSTjAoGBAIghggQheVp+9zrSgJMusuQ3cHuF8Mkg1i1U0Yo0L++JJwwEA3ro81Fr
YKMRQyoVsoBqrpLFFf10hJlnDveN08AAQax007p4H5JtznHG3fKh8kFVYF6e7sY5
c/p1UURL8VJO5dSjQYkSgocJX7YcWZGI01Iy1JDIyKmaS99X8I/c
-----END RSA PRIVATE KEY-----`)

var GoproxyCa, goproxyCaErr = tls.X509KeyPair(CA_CERT, CA_KEY)