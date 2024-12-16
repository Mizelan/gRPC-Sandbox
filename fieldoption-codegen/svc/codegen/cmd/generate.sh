#!/bin/bash
# generate.sh

# 실행 중 오류 발생 시 즉시 스크립트 종료
set -e

# 스크립트의 절대 경로 확인
SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )"

# 프로젝트 루트 디렉토리 설정 (스크립트 기준 3단계 상위)
PROJECT_ROOT="$(cd "$SCRIPT_DIR/../../.." && pwd)"

# 진단 정보 출력
echo "스크립트 디렉토리: $SCRIPT_DIR"
echo "프로젝트 루트 디렉토리: $PROJECT_ROOT"

# protoc-gen-db 플러그인 빌드 및 설치
echo "protoc-gen-db 플러그인 빌드 및 설치..."
cd "$SCRIPT_DIR"
go build -o "$GOPATH/bin/protoc-gen-db" .

# 현재 PATH에 Go 바이너리 디렉토리 추가
export PATH=$PATH:$GOPATH/bin

# proto 파일 컴파일 (사용자 정의 플러그인 사용)
echo "Proto 파일 컴파일 중..."
protoc \
    --proto_path="$PROJECT_ROOT" \
    --db_out="$PROJECT_ROOT/gen" \
    --db_opt=paths=source_relative \
    --go_out="$PROJECT_ROOT/gen" \
    --go_opt=paths=source_relative \
    --go-grpc_out="$PROJECT_ROOT" \
    --go-grpc_opt=paths=source_relative \
    "$PROJECT_ROOT/proto/v1/entity.proto"

echo "코드 생성 완료!"