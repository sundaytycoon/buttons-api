# buttons-api

# 설계 UMLs

[UMLs.md](./UMLs.md)

# gettings started

```
go run cmd/main.go server start # 5001 grpc - 5002 http
```

# package structue?
```
  📂 git@sundaytycoon/buttons-api
   ┣📂 cmd
   ┃ ┗ 📜 main.go # [server, entd]가 있음
   ┣📂 doc
   ┃ ┗ 📂 OpenAPI # swagger.json들어오는곳/ swagger-ui 다운받는 곳
   ┣📂 edge # 비지니스로직 상세하게 타지 않으면서, 외부 서버랑 최종적으로 통신하는 코드들
   ┣📂 ent # entd.go 의 generated된 파일들 한 뭉텅이
   ┣📂 gen/go/buttons/api/v1 # protobuf로 generated된 golang파일
   ┣📂 internal # 내부 비지니스로직 담는곳 # 똥같은 코드들 두는 곳
   ┣📂 pkg
   ┣📂 proto # protobuf 파일들이 있음.
```

``` makefile
make generate-docs # 요거하면 코드 dependency graph가 바뀌어요
```

![code dependency graph](./doc/_images/godepgraph.png)



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

# scripts

### protobuf generating

``` bash
make protosetup # For set up protobuf using `buf`

make protogen # For generating protobuf
```

# References

- git: https://github.com/iDevoid/stygis
- meidum: https://medium.com/@matiasvarela/hexagonal-architecture-in-go-cfd4e436faa3
- youtube: https://www.youtube.com/playlist?list=PLGl1Jc8ErU1w27y8-7Gdcloy1tHO7NriL

