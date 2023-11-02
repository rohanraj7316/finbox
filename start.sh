# # local env creds
export POSTGRESQL_HOST=localhost
export POSTGRESQL_USER=postgres
export POSTGRESQL_PASSWORD=password
export POSTGRESQL_DB_NAME=go_stock_pro
# dev env creds
# export POSTGRESQL_HOST=arjuna.db.elephantsql.com
# export POSTGRESQL_USER=vblgruie
# export POSTGRESQL_PASSWORD=St6Ehn769xwBg3GIKUfpQJ3tV2USxFwv
# export POSTGRESQL_DB_NAME=vblgruie
export POSTGRESQL_LOG_TRACE=false
export POSTGRESQL_PORT=5432
export POSTGRESQL_MAX_OPEN_CONN=5
export POSTGRESQL_MAX_IDLE_CONN=1
export POSTGRESQL_MAX_CONNECTIONS_LIFE_TIME=5s
export POSTGRESQL_CONN_MAX_IDLE_TIME=5s

# postgres://vblgruie:St6Ehn769xwBg3GIKUfpQJ3tV2USxFwv@arjuna.db.elephantsql.com/vblgruie

export REDIS_HOST=redis-12676.c290.ap-northeast-1-2.ec2.cloud.redislabs.com
export REDIS_PORT=12676
export REDIS_AUTH=pKrXx0EUTuF0evkfiXkHSC9RFgnEUIsE
export REDIS_PREFIX=auth_
export REDIS_POOL_SIZE=1
export REDIS_MAX_RETRIES=2

# redis-12676.c290.ap-northeast-1-2.ec2.cloud.redislabs.com:12676

export AUTH_PUBLIC_KEY=MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAw6qD5c0ZkP9Fy2s4GvmHxCkJHEwjVZJMehotDFe7oVy4c9GQdfsTkIdxv85434e3YWfo6oqnVk5k+7I3RAKHw+Xzb81j/HtnY3RUqLSvyyI4CILv95mlGt5OAKD6NGk5RkcA+qDX1Gb/WI9ZVATzrIqIPUQbxluJurKcKWMEF/AJa6kbeiVChf64EN6cpKH+wVUImn6uXx7b+mgC3kVwiDNYSmv1Ouxx+Nc10ZUSbrcdlRDkPMb7RnHOSNoOu6Q/OX+CO64/8h2dyxBaKcgYgzp+Mzph/YdDNLPdOvr6dwoWIwZoRSqlUH4tC/rGRgZey0DEMRwm2Oki3kovURNkzQIDAQAB
export AUTH_PRIVATE_KEY=MIIEvwIBADANBgkqhkiG9w0BAQEFAASCBKkwggSlAgEAAoIBAQDDqoPlzRmQ/0XLazga+YfEKQkcTCNVkkx6Gi0MV7uhXLhz0ZB1+xOQh3G/znjfh7dhZ+jqiqdWTmT7sjdEAofD5fNvzWP8e2djdFSotK/LIjgIgu/3maUa3k4AoPo0aTlGRwD6oNfUZv9Yj1lUBPOsiog9RBvGW4m6spwpYwQX8AlrqRt6JUKF/rgQ3pykof7BVQiafq5fHtv6aALeRXCIM1hKa/U67HH41zXRlRJutx2VEOQ8xvtGcc5I2g67pD85f4I7rj/yHZ3LEFopyBiDOn4zOmH9h0M0s906+vp3ChYjBmhFKqVQfi0L+sZGBl7LQMQxHCbY6SLeSi9RE2TNAgMBAAECggEARbMsqAQvOs0MXkGT61v/JnvdsHRY0+c5EffYG0D3aCxA0pUQyk5FsmLlPMe+nBXa17ptqHr60A49/8EE2dey6DA2TxnRp8OH2VA7xhsTUh+e6T1HYSKcw1z0WVn+twPSsLujWCRKrfGAvVnXHsxixxBUJsrnv/mkrtpoMYABmcq68/iRMZoIjkWpz6wAGN1rwyFrvVZfSJkD5WSjpaQftigW2Vjy4DaK8yPh3NhPi3Exvd/NvdYZKEgyXKzd6DciiEv4YjvtG8OmKexxr+Pu2nVJ2NuyzhILJmUAdWYDjK5vYCPNl75HqECVZOmmL8qOnZx6ZXOHdw358EiaSOaeewKBgQDmR6/IDW0oeTcOeNqIHAzQTPXd4HUTste3uH9X4HrK3gQesa6IYGgMV92ZhXV0Be3tOx4DqMVNuwsr/NSblomVfxdbLpFFHIa5L37lj2Qt1lh7aRoH2HHmhwuL9soUXfcKh405Gs4ooY3CtzAKJ3WnusuYiCaGXPvbG1CrFExdNwKBgQDZhR/aIvd+KHR5xi8DdLZxzi80TWvBxupy0+lTT9yKD56WUslSEAUPteClpgvBb78amDpd78ZdTHybkL8SKc1wk1ugsnTOUSMjxl29bwUEStvbFMOUltJLFwwZ/qswYg0mkNF3dNd++Y7ifO/aSWZxfClw9szmslV6AO+buG7wGwKBgQCWw5pPVJMB65JjNDaG5C5zdd+GuzyLgAyBaRZeV0ataPuziMrm2I9mfWRE7b5/Dp6+MXUuGiHLCUmALBCaKM29Ba8p8GwFnm6J8ZGYA/Annmzhp2b2efvgXMvvf8y/1uE1kJeiKm2M5nBkagDWtQzvmnlPTNxEHb56rgB7cVepswKBgQDWVklWp1H8rFxpihHVc/I7HhKBlTBzV1C6KXomr/D+0flQ4u6hwEtcebmNQJsg1r4WtJ3+5kAuuymmGPFOMobPQUTuiipzCpx3qgXAnl7xJdqSfFlkV2GMwH8aNkn1eQjAUb812Jpn8f0LzFSEYgNiHBixyV+Ki3uSps10Qj5+tQKBgQDRNSiQH1nu2fxgcnfDCscSx6Gi6dKnjBdbT73YmMunBBNMvBhhE3yYK05h6fR+CxqZQ9OrlSFjHYI9t06forZLVskBB+ELcOssedtXDO+303tm5etd6SPLqymliQ9E6ZYDt/ojRhI/P6WD6obAcEp8gJ5gvp2GED96fYPF2rRdSQ==

export INFLUXDB_URL=https://us-east-1-1.aws.cloud2.influxdata.com
export INFLUXDB_TOKEN=FamgEk3cbrgg2LKepAB48HBiDUwAVsQ4YIXyukocn_j4m3eBbyMa6SJVozncK9Q0_AC9D9Fog8XaMr2kgwE_Aw==

go run main.go