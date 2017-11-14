# Request time challenge 

Vamos desenvolver uma aplicação chamada CRGO. 

O CRGO é uma aplicação para verificar quanto tempo demora para fazer pedidos
a uma ou mais URLs, também informa quantos `BYTES` o `HTML` tem na sua resposta. 

---

### Helpers
       
- Packages used:

    * [time](https://golang.org/pkg/time)
    * [net/http](https://golang.org/pkg/net/http)
    * [fmt](https://golang.org/pkg/fmt)
    * [io/ioutil](https://golang.org/pkg/io/ioutil)
    * [runtime](https://golang.org/pkg/runtime)
      
- References:
    
    * [Variables](https://gobyexample.com/variables)
    * [For](https://gobyexample.com/for)
    * [If-else](https://gobyexample.com/if-else)
    * [Errors](https://gobyexample.com/errors)
    * [Arrays](https://gobyexample.com/arrays)
    * [Range](https://gobyexample.com/range)
    * [Gorotines](https://gobyexample.com/goroutines)
    * [Channels](https://gobyexample.com/channels)
---

### Step 1:
    
>  É preciso criar uma função que irá receber uma string url e fazer um pedido http para saber quanto tempo esse pedido demora a responder e deve retornar uma string.
   
Input: 
    
    "http://www.terra.com.br/" 
    
Output: 

    0.98s
    
---

### Step 2: 

> Vamos formatar a mensagem de output para ficar mais legível.

Input: 
       
    "http://www.terra.com.br/"
    
Output: 

    [0.98s] elapsed time for request [http://www.terra.com.br/]
 
---

### Step 3:

> O mesmo pedido agora deve retornar o `length` dos `BYTES` que a resposta do body tem e adicionar no output.
    
Input: 
       
    "http://www.terra.com.br/"
    
Output: 

    [0.98s] elapsed time for request [http://www.terra.com.br/] with [262425] bytes
 
---

### Step 4: 

> Agora precisamos fazer com que nossa aplicação possa receber um ou mais urls para fazer 
mais pedidos.

Input:   

    http://www.terra.com.br/ http://www.google.com/ http://www.facebook.com/

Output: 

    [0.98s] elapsed time for request [http://www.terra.com.br/] with [262425] bytes
    [0.14s] elapsed time for request [http://www.google.com/] with [11546] bytes
    [0.86s] elapsed time for request [http://www.facebook.com/] with [88051] bytes
    
---

### Step 5: 
    
> Vamos agora mostrar o tempo total de execução do começo da execução ate o final.
    
Input:  
    
    http://www.google.com/ http://www.facebook.com/ http://www.terra.com.br/
    
Output: 

    [0.98s] elapsed time for request [http://www.terra.com.br/] with [262425] bytes
    [0.14s] elapsed time for request [http://www.google.com/] with [11546] bytes
    [0.86s] elapsed time for request [http://www.facebook.com/] with [88051] bytes
    [3.33s] elapsed time.
    
---

### Step 6: 
    
> Agora nossa aplicação deve nos mostrar os resultados em ordem na qual chegam, isso é o primeiro a chegar 
> estará no topo da nossa lista, o segundo a chegar sera o segundo e assim sucessivamente.

Input: 

    http://www.google.com/ http://www.facebook.com/ http://www.terra.com.br/
    
Output: 

    [0.14s] elapsed time for request [http://www.google.com/] with [11546] 
    [0.86s] elapsed time for request [http://www.facebook.com/] with [88051] 
    [0.98s] elapsed time for request [http://www.terra.com.br/] with [262425] 
    [3.33s] elapsed time.
    
---

### Step 7: 

> Agora que o CRGO está quase pronto vamos fazer com que a ele possa receber as urls com parametros 
> para que possamos usar diretamente no terminal

Input: 

    $ ./crgo http://www.google.com/ http://www.facebook.com/ http://www.terra.com.br/

Output:

    [0.14s] elapsed time for request [http://www.google.com/] with [11546] 
    [0.86s] elapsed time for request [http://www.facebook.com/] with [88051] 
    [0.98s] elapsed time for request [http://www.terra.com.br/] with [262425] 
    [3.33s] elapsed time.
    
---

### Step 8: 

> Para finalizar vamos melhorar a performance do CRGO, vamos fazer com que o tempo de execução total seja 
> sempre o do mais demorado, isso é se um request para a url `X` demorar mais do que a `Y` a o tempo do 
> programa deve ser aproximado com o do `X` mais o tempo de processamento, os pedidos devem rodar em **PARALELO**. 

Input: 
 
     $ ./crgo http://www.google.com/ http://www.facebook.com/ http://www.terra.com.br/
     
Output 

    [0.14s] elapsed time for request [http://www.google.com/] with [11546] 
    [0.86s] elapsed time for request [http://www.facebook.com/] with [88051] 
    [0.98s] elapsed time for request [http://www.terra.com.br/] with [262425] 
    [1.77s] elapsed time.
    
---

### Step 9:

> Agora podemos usar um Channel para a resposta do nosso projeto, onde todos os output fiquem centralizados 
> passando por um channel. 

Input: 

    $ ./crgo http://www.google.com/ http://www.facebook.com/ http://www.terra.com.br/
  
Output 

    [0.14s] elapsed time for request [http://www.google.com/] with [11546] 
    [0.86s] elapsed time for request [http://www.facebook.com/] with [88051] 
    [0.98s] elapsed time for request [http://www.terra.com.br/] with [262425] 
    [1.77s] elapsed time.

