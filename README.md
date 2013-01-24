## Go Tour Korean translation source and tool

Go Tour(Korean) Web: [http://go-tour-kr.appspot.com](http://go-tour-kr.appspot.com)  
Go Tour(Korean) GitHub: [http://github.com/golanger/go-tour-kr](http://github.com/golanger/go-tour-kr)  

tour.article 파일은 Go의 present 패키지를 이용해서 html로 변환해줄 수 있는 원본 파일입니다. 이 파일의 문법은 [present package](http://godoc.org/code.google.com/p/go.talks/pkg/present)를 참고해주세요.


### 1. Step 1: tour.article 파일 번역

    tour.article.orig 파일은 tour.golang.org의 원본 파일입니다. tour.article을 번역해서 pull request 보내주시면 됩니다.

### 2. Step 2: index.html로 변환

    $ go run conv.go

    위 명령을 실행하시면 번역된 tour.artcle 파일을 index.html로 변환시켜 줍니다.
    이 파일을 Go Tour App Engine 소스의 static 디렉토리에 복사해서 앱 엔진 서버에 올리면 번역된 내용이 적용됩니다.
