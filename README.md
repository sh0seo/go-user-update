# Golang을 이용한 DCCafe 사용자 정보 업데이트

- DigiCAP에서 운영 중인 Cafe의 사용자 정보를 자동 갱신하는 goland 코드
- 경영지원실에서 잘못 등록하는 사용자 정보를 갱신하기 위해 사용

## 사용방법

- main.go에 name, cardNumber에 변경해야하는 사용자의 이메일과 변경하려는 카드 번호를 입력

```
// input
email := "rfid_tester@email.com"
// input
cardNumber := "123456789"
```

- 변경한 코드를 master에 push 하면 자동으로 갱신이 됨

## 정상 처리 확인 방법

- Action -> ActionFlows 로그에서 아래와 같이 나오면 성공
```
Run go run main.go
DCCafe User RFID Update
find user 김재경
Get new RFID dc5ffb37f6e0f1df638853a9cd0549f950caa594f3529ca060e7a2fae0808b8d
Update User: 김재경, CardNumber: 3255213734
```

- 아래와 같이 나오면 실패

```
Run go run main.go
DCCafe User RFID Update
find user rfid_tester
[Error] Exist Card Number:123456789
```
