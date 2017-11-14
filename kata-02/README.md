# Duplicate files challenge

Hoje em dia é muito comum termos muitas fotos armazenadas em nosso HD, fazemos backup
e mantemos sempre uma copia em vários lugares.

Imagine que você tem duas vezes a mesma imagem apenas com nome diferente, o conteúdo é o mesmo
porém com nomes diferentes.
Vamos fazer um programa que resolva isso, ele verificará o conteúdo do arquivo caso tenha em duplicado 
irá mover de pasta.

--- 

### Helpers
       
- Packages used:

    * [path/filepath](https://goglang.org/pkg/path/filepath)
    * [os](https://goglang.org/pkg/os)
    * [fmt](https://goglang.org/pkg/fmt)
    * [crypto/md5](https://goglang.org/pkg/crypto/md5)
    * [io/ioutil](https://goglang.org/pkg/io/ioutil)
    * [log](https://goglang.org/pkg/log)
    * [path](https://goglang.org/pkg/path)
      
- References:
    
    * [Variables](https://gobyexample.com/variables)
    * [If-else](https://gobyexample.com/if-else)
    * [For](https://gobyexample.com/for)
    * [Struct](https://gobyexample.com/structs)
    * [Arrays](https://gobyexample.com/arrays)
    * [Slices](https://gobyexample.com/slices)
    * [Range](https://gobyexample.com/range)
    * [Methods](https://gobyexample.com/methods)
    * [Closures](https://gobyexample.com/closures)
    
---

###  Step 1 

> Você precisa listar todos arquivos de uma determinada pasta. 
> Faça uma função onde liste todos os arquivos de uma pasta

Input: 

    fotos_dir
    
Output:

    GS2iiTz8_400x400.png
    Go-aviator.sh.png
    fix-it.jpeg
    fixing-code.jpeg
    foca no code.jpg
    
    
> **Tip**: [filepath](https://golang.org/pkg/path/filepath/#Walk) 

OBS.: 
    .          

---

### Step 2 

> Agora deve criar uma função na qual irá ler o conteúdo dos arquivos e gerar uma HASH única para cada arquivo. 

Input: 

    fotos_dir

Output: 

    [24a35fdfb952e71574ce26cdb0fb6464] GS2iiTz8_400x400.png
    [ab53300f7fe15bbc9dd7d4067eb306bb] Go-aviator.sh.png
    [f239c089076fe15c61e401f7fb6bacc0] fix-it.jpeg
    [f239c089076fe15c61e401f7fb6bacc0] fixing-code.jpeg
    [e242a545cebc29e0deea89862aada894] foca no code.jpg

**tips**: pode usar : [MD5](https://golang.org/pkg/crypto/md5/#pkg-examples)

---

### Step 3

> Agora você precisa mostrar quais as imagens possuem o mesmo conteúdo

Input: 

    fotos_dir

Output:     
    
    File [gopher-keith-57.png] is unique.
    File [playing.jpg] is unique.
    File [sing.jpg] is unique.
    File [strong.png] is unique.
    File [Go-aviator.sh.png] is unique.
    Exists [2] files with [f239c089076fe15c61e401f7fb6bacc0] hash
             0 - [fix-it.jpeg].
             1 - [fixing-code.jpeg].
    File [global.jpeg] is unique.


