# Golang을 이용한 DCCafe 사용자 정보 업데이트

- DigiCAP에서 운영 중인 Cafe의 사용자 정보를 자동 갱신하는 goland 코드
- 경영지원실에서 잘못 등록하는 사용자 정보를 갱신하기 위해 사용

## 사용방법

- main.go에 name, cardNumber에 변경해야하는 사용자와 변경하려는 카드 번호를 입력

```
	// input
	name := "rfid_tester"
	// input
	cardNumber := "123456789"
```

- 변경한 코드를 master에 push 하면 자동으로 갱신이 됨
