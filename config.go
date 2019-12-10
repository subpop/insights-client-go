package main

type config struct {
	CertFile string
	KeyFile  string
}

func defaultConfig() *config {
	return &config{
		CertFile: "/etc/pki/consumer/cert.pem",
		KeyFile:  "/etc/pki/consumer/key.pem",
	}
}
