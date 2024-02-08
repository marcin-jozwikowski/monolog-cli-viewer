package data

func GetTestData() []string {
	return []string{
		`{"message":"Some test message","context":{"user":{"id":1},"session":{"id":"bq2fk4i3nhkgbj4eua964g5r63"}},"level":"NOTICE","channel":"default","timestamp":"1699540146122"}`,
		`{"message":"Checking support on authenticator.","context":{"firewall_name":"main","authenticator":"App\\Security\\AppAuthenticator"},"level":100,"level_name":"DEBUG","channel":"security","datetime":"2023-11-14T00:37:26.623539+02:00","extra":{}}`,
		`2023-10-23 11:03:16: [9a4e77e9afa8] [ERROR] [Whatever] Login Error`,
		`==> some/path/to\file.log <==`,
		`[2023-10-23T11:07:47.038324+00:00] default.INFO: User logged in {"user":{"id":"54767261-98c6-4a57-9064-0d35fd06d1fc"}} []`,
		`[2023-12-29T10:26:40.537772+00:00] request.INFO: Matched route "api_login". {"route":"api_login","route_parameters":{"_route":"api_login","_controller":"App\\User\\Infrastructure\\Controller\\ApiLoginController::index"},"request_uri":"http://localhost/api/v1/login","method":"POST"} []`,
	}
}
