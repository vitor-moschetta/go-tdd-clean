## Indice

Desenvolvimento orientado a testes (TDD) com Casos de uso (UseCase - Clean Architecture) bem definidos.

01. Input basico e UseCase basico retornando true
02. Input com Validação
03. Agrega Entidade e Interface do Repositorio. 
    - Acoplamento, injeção de dependencia e inversão de controle.     
04. Agregamos um Repositorio em Memoria (mock)    
05. Agregamos Validação de Entidade. Aqui podemos falar sobre validação de entidade e validação de input
    - Porque não validamos somente a entidade?         
        - Porque precisamos fornecer feedback para o usuário sobre os seus dados de entrada (Isso dará abertura para entrarmos com o conceito de Notification Pattern)
        - Porque a validação de entidade pode envolver regras mais complexas que a validação de input.
        - Fail Fast
06. Agregamos propriedade Price na Entidade
    - Fazemos uma introdução sobre Notification Pattern: Na primeira requisição o usuário precisa saber todas as falhas que ocorreram.
07. Agregamos Notification Pattern com Output do tipo []errors
08. Agregamos Notification Pattern com Output do tipo struct
    - Encapsulamento: usar apenas codigos de domínio válidos
09. Piramide de testes: Testes de unidade e testes de fluxo de caso de uso
10. Organização para test coverage
    go test ./... -coverprofile=coverage.out
    go tool cover -html=coverage.out
11. Adiciona QueryInput para buscar produtos por intervalo de preço
12. Adiciona Infraestrutura Web (API REST)

## Postman
O projeto 12 possui infraestrutura web (API REST) e pode ser testado com o Postman. A collection a ser importada está na pasta `postman`.

```bash
curl -X GET "http://localhost:8080/api/v1/products?min_price=0&max_price=200" -H "accept: application/json"
```


## Índice (ES)

Desarrollo orientado a pruebas (TDD) con Casos de uso (UseCase - Clean Architecture) bien definidos.

01. Input básica y UseCase básico que devuelve true
02. Input con validación
03. Agrega Entidad e Interface del Repositorio.
    - Acoplamiento, inyección de dependencia e inversión de control.
    - Aquí dará error de referencia porque no tenemos una implementación del repositorio.
04. Agregamos un Repositorio en Memoria (mock)
    - Tendremos error de referencia circular entre el archivo de prueba y el archivo de mock. Necesario crear paquete para archivo de prueba.
05. Agregamos Validación de Entidad. Aquí podemos hablar sobre validación de entidad y validación de input
    - ¿Por qué no validamos solo la entidad?
        - Porque necesitamos proporcionar comentarios al usuario sobre sus datos de entrada (Esto abrirá el concepto de Notification Pattern)
        - Porque la validación de entidad puede implicar reglas más complejas que la validación de entrada.
        - Fail Fast
06. Agregamos propiedad Price en la Entidad
    - Hacemos una introducción sobre Notification Pattern: En la primera solicitud, el usuario necesita saber todas las fallas que ocurrieron.
07. Agregamos Notification Pattern con Output del tipo []errors
08. Agregamos Notification Pattern con Output del tipo struct
    - Encapsulamiento: usar solo códigos de dominio válidos
09. Pirámide de pruebas: Pruebas de unidad y pruebas de flujo de caso de uso
10. Organización para test coverage
    go test ./... -coverprofile=coverage.out
    go tool cover -html=coverage.out
11. Agrega QueryInput para buscar productos por intervalo de precio
12. Agrega Infraestructura Web (API REST)

## Postman (ES)
El proyecto 12 tiene infraestructura web (API REST) y se puede probar con Postman. La colección a importar está en la carpeta `postman`.

```bash
curl -X GET "http://localhost:8080/api/v1/products?min_price=0&max_price=200" -H "accept: application/json"
```