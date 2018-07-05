# DMS-Migrates-to-Go
10개월 전에 파이썬과 Flask를 배우고 나서, Vert.x 기반의 DMS 서버를 Python + Flask + MongoDB 기반으로 [마이그레이션하는 연습](https://github.com/JoMingyu/DMS-Migrates-to-Python)을 해 본 적 있었다. 이번엔 Go를 배웠으니, Go + Echo + MongoDB 조합으로 DMS 서버를 만들어 보자.

## 규칙
- DMS 내부적으로 secret하게 공유되고 있는 Swagger를 기반으로 하되, API spec과 request/response payload는 맘대로 바꾸자.(기능 구현을 위주로)

## Checklist
- [x] 프로젝트 구조 잡기
- [x] MongoDB 데이터베이스 connector 패키지
- [ ] API 로직 작성(인증과 암호화 관련 부분 제외)
    - [ ] 학생 API
        - [x] 계정
        - [ ] 계정 관리
        - [ ] 신청
        - [ ] 신고
    - [ ] 관리자 API
        - [ ] 계정
        - [ ] 계정 관리
        - [ ] 신청 정보
        - [ ] 상벌점 관리
        - [ ] 게시글 관리
        - [ ] 신고 정보 관리
        - [ ] 설문지 관리
    - [ ] 공통 API
        - [ ] JWT 관련
        - [ ] 메타데이터
        - [ ] 게시글
        - [ ] 학교 정보
- [ ] JWT 붙이기
- [ ] 데이터 암호화 붙이기
- [ ] Swagger 붙이기
- [ ] 테스트 코드 작성
- [ ] CI 붙이기
- [ ] 로그 데이터 시각화를 위한 백오피스 구축