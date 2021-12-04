# profile.me-server

## Project structure

1. cmd

   - Golang으로 작성된 프로그램의 앤트리포인트인 main package가 있습니다. main function은 init 작업을 수행한 후 cobra의 root command를 실행합니다. golive에서 실행되는 여러 커맨드는 cmd 아래의 하위 디렉터리로 존재합니다. `ex) golive server-swager => cmd/sever/swagger`

2. handler

   - 외부에서 들어오는 요청(http, gRPC) 등을 차리합니다. 외부에서 들어온 요청을 validation하고, 알맞은 내부 서비스를 호출합니다. 이후 caller에게 응답을 줍니다.
   - 구독하고 있는 메시지(kafka, kinesis, sqs) 등을 처리합니다. 외부에서 가져온 메시지를 validation하고 알맞은 내부 서비스를 호출합니다. 이후 처리한 메시지를 ACK 처리합니다.
   - `handler/websocket/hub.go`
     - 서버가 맺은 웹소켓 커넥션을 관리합니다.
     - session은 웹소켓 커넥션의 정보를 가지고 있습니다.
     - hub에서 관리되는 각 channel에 대해서 세션들이 묶여있습니다.
     - hub > liveStreamID > sessions 로 데이터 상하관계를 가지고 있습니다.

3. api

   - 핸들러에서 다루는 요청의 명세를 다룹니다. gRPC, Open API Specifications의 경우 명세에 따라 코드가 자동으로 생성됩니다.

4. service

   - 비즈니스 로직을 오케스트레이션합니다. 유즈케이스에 따라 command 혹은 event를 인자로 받습니다.

5. repo

   - 객체 저장과 접근을 제공하는 저장소입니다. 저장소는 어그리게이트만 반환하도록 강제합니다.

6. domain

   - `domain/model`은 각 레이어 간에 호환되는 데이터 오브젝트입니다. 예를 들어 `entgo` ORM을 내부에서 사용하는 레포지토리는 다른 레이어로 데이터 오브젝트를 전달할 때 model의 객체로 변환하여 전달해야 합니다. 흔히 DTO라고 일컫는 것과 비슷합니다.
   - `domain/message/command`는 profile.me-server 서비스에서 전달되는 메시지의 일부입니다. 커맨드는 한 행위자로부터 다른 구체적인 행위자에게 전달됩니다. 보내는 행위자는 받는 행위자가 커맨드를 받고 구체적인 작업을 수행하길 기대합니다. API 핸들러로 호출하는 행위는 커맨드를 호출하는 행위와 같습니다. 그래서 이름을 붙일 때, 명령형으로 사용합니다. `ex) GetUser, EditUser` 커맨드는 어떤 일을 수행하기를 바라고 그 수행의 결과를 돌려받기를 원합니다.
   - `domain/message/event`는 행위자가 관심있는 모든 리스너에게 보내는 메시지입니다. `LivestreamCreated` 라는 이벤트를 발행해도 발행하는 행위자는 누가 이 이벤트를 받는지 모릅니다. 이벤트를 보내는 쪽은 받는 쪽의 행위가 성공하는지 실패하는지 관심이 없습니다.

7. adapter

   - 서비스나 저장소에서 외부와 통신하기 위해 사용하는 프로토콜, 데이터베이스 커넥터 등의 구현체입니다.

8. config

   - 애플리케이션에서 전역적으로 사용하는 설정 값을 환경변수로 받습니다.

9. DI direction
   - server => handler (repo => service)
