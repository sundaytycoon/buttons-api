# buttons-api

# μ€κ³ UMLs

[UMLs.md](./UMLs.md)

# gettings started

```
go run cmd/main.go server start # 5001 grpc - 5002 http
```

# package structue?
```
  π git@sundaytycoon/buttons-api
   β£π cmd
   β β π main.go # [server, entd]κ° μμ
   β£π doc
   β β π OpenAPI # swagger.jsonλ€μ΄μ€λκ³³/ swagger-ui λ€μ΄λ°λ κ³³
   β£π edge # λΉμ§λμ€λ‘μ§ μμΈνκ² νμ§ μμΌλ©΄μ, μΈλΆ μλ²λ μ΅μ’μ μΌλ‘ ν΅μ νλ μ½λλ€
   β£π ent # entd.go μ generatedλ νμΌλ€ ν λ­νμ΄
   β£π gen/go/buttons/api/v1 # protobufλ‘ generatedλ golangνμΌ
   β£π internal # λ΄λΆ λΉμ§λμ€λ‘μ§ λ΄λκ³³ # λ₯κ°μ μ½λλ€ λλ κ³³
   β£π pkg
   β£π proto # protobuf νμΌλ€μ΄ μμ.
```

``` makefile
make generate-docs # μκ±°νλ©΄ μ½λ dependency graphκ° λ°λμ΄μ
```

![code dependency graph](./doc/_images/godepgraph.png)



# wiki

### ent.go μ μ©ν μ€ν¬λ¦½νΈ λͺ¨μ

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

### κ°λ° μν κ΄λ ¨ λ¬Έμ

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

