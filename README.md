# profile.me-server

TDD코드를 잘 짜고 싶은건 맞는데, 어떤 코드가 테스트짜기 좋은코드인지 잘 모르겠어서 골라본 코드페턴 hexagonal(핫 한듯해서)

### package structure

1. handler [handler]
2. (interceptor by handler) [middleware]
3. service[do business logic]
4. repository[user, profile, article model]
5. storage[serviceDB, cache, search engine]
- constants => 정적 쿼리 및 정적 모델(DTO, VO)

# References

- git: https://github.com/iDevoid/stygis
- meidum: https://medium.com/@matiasvarela/hexagonal-architecture-in-go-cfd4e436faa3
- youtube: https://www.youtube.com/playlist?list=PLGl1Jc8ErU1w27y8-7Gdcloy1tHO7NriL

# 
- buf 를 이용한 Protobuf 관리

# scripts

### protobuf generating

``` bash
make protosetup # For set up protobuf using `buf`

make protogen # For generating protobuf
```