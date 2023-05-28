## Restaurant menu voting app

Test task for Genesis software engineering school 3.0

### Installation:

 ```bash
  git clone https://github.com/rraatt/BTCrateApi.git
```

Create a .env file and add the following variables:

 ```bash
  API_KEY = <Your coinmarketcap api key>
  SMTP_USERNAME = <Your SMTP username>
  SMTP_PASSWORD = <You SMTP password>
  SMTP_HOST = <Your SMTP host>
  SMTP_PORT = <Your SMTP port>
  SMTP_SENDER = <Your sender email>
```

Create and activate a virtual environment:

 ```bash
  docker-compose up -d --build
```

### Usage:

Access at localhost:

 http://localhost:8080/rate

 http://localhost:8080/subscribe

 http://localhost:8080/sendEmails
