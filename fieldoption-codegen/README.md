# fieldoption 테스트

## 목적
- fieldoption을 proto 속성으로 code generation을 통해 코드를 생성한다.

## 실행 순서
1. codgen 바이너리 빌드를 위해 `entity.pb.go` 생성
    ```shell
    # proto/generate.go의 annotation을 실행
    $ go generate ./...
    ```
1. codegen 실행
    ```shell
    $ cd ./svc/codegen/cmd
    $ ./generate.sh
    ```
1. 생성된 코드 확인
   - `gen/proto/v1/entity.db.go` 에 생성됨