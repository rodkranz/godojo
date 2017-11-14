
---
### Step 1:

Primeira coisa é ir a **API** do github para a estrutura dos dados que vamos usar. 

O link é `https://api.github.com/users/{:username}`

Link: [GitHub](https://api.github.com/users/rodkranz)

---
### Step 2: 

Devemos criar uma estrutura com os dados que vamos precisar.  

**UserGithub**:

| Name         | Type   |
|--------------|--------|
| login        | string |    
| html_url     | string |    
| avatar_url   | string |    
| name         | string |    
| company      | string |    
        
---
### Step 3: 

Agora vamos buscar os dados da **API** e popular na estrutura que criamos, precisamos criar uma função com a assinatura `func fetchUserInfoFromGithub(username string) (gu GithubUser, err error)`
para que possamos testar.
 
Depois de executar a função para testar fazemos um `Printf` da `struct` com formato `%#v`.

Output esperado é:
```bash:
go run main.go
    main.UserGithub{Login:"rodkranz", HtmlUrl:"https://github.com/rodkranz", Name:"Rodrigo Lopes", AvatarURL:"https://avatars2.githubusercontent.com/u/16897636?v=4", Company:"OLX", PublicRepos:34, PublicGists:20, Followers:23, Following:1} 
```

> **Dica**
> * Para mostrar no mesmo formato você pode usar o `fmt` com o formato `Printf("%#v\n", ...)`. [Documentation](https://golang.org/pkg/fmt/#Printf)
> * Você vai precisa usar o package `net/http` para fazer o fetch dos dados. [Documentation](https://golang.org/pkg/net/http/#example_Get)
> * Você tambem vai precisar do package `encoding/json` para fazer o decode do json para a `struct`. [Documentation](https://golang.org/pkg/encoding/json/#NewDecoder)
> * Deve retornar erro se o status code for diferente de 200.
> * Se o status code for 404 deve retornar a mensagem ``user not found``.

--- 
### Step 4:

Já temos todas as informações que precisamos do `user` do github, agora vamos ver como podemos fazer para enviar uma nova mensagem no `Slack`.

Começamos criando duas `struct` para a mensagem, Uma `SlackMessage` e outra `Attachment` que devem ter as seguintes propriedades:

1 **Attachment**: 

| Name        | Type    |
|-------------|---------|
| Color       |  string |
| AuthorName  |  string |
| AuthorLink  |  string |
| AuthorIcon  |  string |
| Title       |  string |
| TitleLink   |  string |
| Text        |  string |

2 **SlackMessage**: 

| Name        | Type         |
|-------------|--------------|
| Channel     | string       |        
| Username    | string       |        
| IconEmoji   | string       |        
| Attachments | []Attachment |  

Para mais informações [Documentation Slack Attaching](https://api.slack.com/docs/message-attachments)

---
### Step 5:

Vamos criar um metodo que ira receber a nossa primeira `struct`(`UserGithub`) e retorna a segunda `struct`(`SlackMessage`) com os dados populados, o metodo tera a seguinte assinatura `func hydrateMessage(github UserGithub) SlackMessage`:

Você deve associar as informações:

| SlackMessage Fields | Values     |
|---------------------|------------|
|  Channel            | "#rodrigo" |               
|  Username           | "Tomilio"  |
|  IconEmoji          | ":gopher:" |              
|  Attachments        | Attachment |       


| Attachment Fields   | Values     |
|---------------------|------------|
|  Color              | "#36a64f"      
|  AuthorIcon         | github.AvatarURL        
|  AuthorName         | github.Name         
|  AuthorLink         | github.HtmlUrl         
|  TitleLink          | github.HtmlUrl         
|  Title              | github.Login     
|  Text               | "{github.Name} is working at {github.Company} company."  


---
### Step 6

Temos quase tudo que precisamos para fazer um post no `Slack` com a nossa mensagem, antes de fazer o post devemos converter nossa `struct` para um formato `json` no qual a **API** do `slack` conhece,
para isso vamos fazer um **bind** na nossa estrutura `SlackMessage` de um novo metodo com a seguinte assinatura `func Bytes() []bytes`.

O metodo chamado **Bytes** deve returnar um slices de bytes 

> **Dica**
> * Associar metodos a estrutura [GoByExample Methods](https://gobyexample.com/methods)
> * Transformar `struct` em `json` você deve usar o `json.Marshal`. [Documentation](https://golang.org/pkg/encoding/json/#Marshal)  
> * Slices em GO [GoByExample Slices](https://gobyexample.com/slices) 

---
### Step 7

Agora temos tudo que precisamos para fazer o post na **API** do `slack`, vamos criar um metodo `func sendMessageToSlack(message SlackMessage) error`,
que vai fazer o post na `API` do `slack`: `https://hooks.slack.com/services/XXXXXXXXX/XXXXXXXXX/000000000000000000`.

> **Dicas**
> * Você vai precisar de fazer um post para isso use o metodo `http.Post` do package `net/http`. [documentation](https://golang.org/pkg/net/http/#Client.Post)
> * Vai precisar fazer um *stream* de buffer para isso use `bytes.NewBuffer` do package `bytes`. [documentation](https://golang.org/pkg/bytes/#NewBuffer)

---
### Step 8

Agora Vamos rodar nossa aplicação e ver a magica acontecer.  
```bash
$ go run mais.go
```
