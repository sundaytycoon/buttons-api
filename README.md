# profile.me-server


### package structure

1. handler [handler]
2. (interceptor by handler) [middleware]
3. service[do business logic]
4. repository[user, profile, article model]
5. storage[serviceDB, cache, search engine]

- constants => 정적 쿼리 및 정적 모델(DTO, VO)