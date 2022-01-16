# profile.me-server

# wiki

### ent.go 유용한 스크립트 모음

- [Link: Official Refereneces](https://entgo.io/docs/getting-started)

``` markdown
# add model
go run entgo.io/ent/cmd/ent init User

# sync schema
### First, should to modify ent/schema/*.go before synchronizing schema
vi ./ent/schema/*.go # step 1
go generate ./ent # step 2
go run cmd/main.go ent migration # step 3


```

### 개발 셋팅 관련 문서

- [DEVELOPMENT.md](https://github.com/sundaytycoon/buttons-api/blob/main/DEVELOPMENT.md)


### package structure

TDD코드를 잘 짜고 싶은건 맞는데, 어떤 코드가 테스트짜기 좋은코드인지 잘 모르겠어서 골라본 코드페턴 hexagonal(핫 한듯해서)

``` makefile
make generate-docs # 요거하면 토드 dependency graph가 바뀌어요
```

![code dependency graph](./doc/_images/godepgraph.png)


1. infrastructure
2. handler [handler]
3. (interceptor by handler) [middleware]
4. service[do business logic]
5. repository[user, profile, article model]
6. storage[serviceDB, cache, search engine]
- constants => 정적 쿼리 및 정적 모델(DTO, VO)

# References

- git: https://github.com/iDevoid/stygis
- meidum: https://medium.com/@matiasvarela/hexagonal-architecture-in-go-cfd4e436faa3
- youtube: https://www.youtube.com/playlist?list=PLGl1Jc8ErU1w27y8-7Gdcloy1tHO7NriL

# scripts

### protobuf generating

``` bash
make protosetup # For set up protobuf using `buf`

make protogen # For generating protobuf
```
