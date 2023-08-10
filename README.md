## Indice

Desenvolvimento orientado a testes (TDD) com Casos de uso (UseCase - Clean Architecture) bem definidos.

01. Estrutura inicial
    - Input basico 
    - UseCase vazio, retornando (output) "error" 
02. Validação de Input do Usuário
03. Persistência de dados
    Para persistir o dado precisamos:
        - Enriquecer os dados (Entidade)
        - Ter um local para armazenar os dados (Repositorio)
            - Não queremos nos preocupar com a implementação do repositorio (inversão de controle)
            - Não queremos nos preocupar com a instanciação do repositorio (injeção de dependencia)
            - Queremos testar o caso de uso sem depender de um banco de dados (mock do repositorio)
04. Notification Pattern
    - Agregamos propriedade Price na Entidade e Input
    - Como podemos notificar todos os erros de validação em uma única requisição?
        - R: Adicionamos um "error" customizado nas validações. Alternativas: []string, []error, struct output, etc..
05. Validação de Entidade
    - Agregamos Data de Cadastro na Entidade
    - Agregamos validação de Entidade
    - Movemos a criação e validação da entidade para seu próprio contexto (DDD, entidades anemicas vs entidades ricas, testes unitarios vs testes de fluxo, piramide de testes)
    - Porque ainda temos que validar o input do usuário? Não poderíamos apenas validar a entidade?
        - R: Porque a entidade pode conter regras internas, e o input é mais específico para a validação dos dados de entrada do usuário  
        - R: Fail Fast
06. Exportar dados de dominio
    - Agregar Output no UseCase (além do error) para retornar os dados de dominio (Entidade)    
07. Exportar Codigos de dominio
    - Como a camada de apresentação poderá mapear e informar corretamente os erros que ocorreram no dominio?    
        - R: Criamos um output customizado (struct) para o dominio , com seus próprios códigos de erro
08. Agregar novo caso de uso
    - QueryInput para buscar produtos por intervalo de preço
    - Agrega query interface repository
    - Agrega query repository mock
    - Agrega seed de dados
09. Organização em contextos (bounded contexts)
    - Adiciona novo contexto (Category)    
    - Shared Context
10. UseCase chamando outro UseCase. UseCase de uma entidade chamando Repository de outra entidade.
    - UseCase de cadastro de produto chama UseCase de cadastro de categoria (Mediator Pattern)
    - UseCase de cadastro de produto chama Repository de categoria (Mediator Pattern)
11. Outros conceitos importantes
    - Unit of Work
    - Domain Events
99. Adiciona Infraestrutura Web (API REST)

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