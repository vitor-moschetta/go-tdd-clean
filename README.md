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
20. Adiciona Infraestrutura Web (API REST)


### Outros temas de domínio a serem abordados
- Comunicação entre casos de uso (Mediator)
- Container de acesso aos repositórios (Repository Manager)
- Agrupar transações de escrita em uma única transação (Unit of Work)
- Eventos de Domínio (Domain Events)

### Outros temas de API a serem abordados
- Versionamento de API
- Pagination and Sorting
- Logging
- Error Handling
- Middlewares
- Hooks
- Cache
    - In-Memory
    - Distributed
- Tracing   
- Timeout e Retries
- Health Check
- Rate Limit
- Circuit Breaker
- Testes de integração
- Testes de carga
- Documentação de API


## Postman
O projeto 12 possui infraestrutura web (API REST) e pode ser testado com o Postman. A collection a ser importada está na pasta `postman`.

```bash
curl -X GET "http://localhost:8080/api/v1/products?min_price=0&max_price=200" -H "accept: application/json"
```


## Índice (ES)

01. Estructura inicial
    - Input basico 
    - UseCase vacio, retornando (output) "error"
02. Validación de Input del Usuario
03. Persistencia de datos
    Para persistir el dato necesitamos:
        - Enriquecer los datos (Entidad)
        - Tener un lugar para almacenar los datos (Repositorio)
            - No queremos preocuparnos por la implementación del repositorio (inversión de control)
            - No queremos preocuparnos por la instanciación del repositorio (inyección de dependencia)
            - Queremos probar el caso de uso sin depender de una base de datos (mock del repositorio)
04. Notification Pattern
    - Agregamos propiedad Price en la Entidad y Input
    - ¿Cómo podemos notificar todos los errores de validación en una única solicitud?
        - R: Agregamos un "error" personalizado en las validaciones. Alternativas: []string, []error, struct output, etc..
05. Validación de Entidad
    - Agregamos Data de Cadastro en la Entidad
    - Agregamos validación de Entidad
    - Movemos la creación y validación de la entidad a su propio contexto (DDD, entidades anémicas vs entidades ricas, pruebas unitarias vs pruebas de flujo, pirámide de pruebas)
    - ¿Por qué todavía tenemos que validar la entrada del usuario? ¿No podríamos simplemente validar la entidad?
        - R: Porque la entidad puede contener reglas internas, y la entrada es más específica para la validación de los datos de entrada del usuario
        - R: Fail Fast
06. Exportar datos de dominio
    - Agregar Output en UseCase (además del error) para devolver los datos de dominio (Entidad)
07. Exportar Códigos de dominio
    - ¿Cómo podrá la capa de presentación asignar e informar correctamente los errores que ocurrieron en el dominio?
        - R: Creamos una salida personalizada (struct) para el dominio, con sus propios códigos de error
08. Agregar nuevo caso de uso
    - QueryInput para buscar productos por intervalo de precio
    - Agrega consulta de interfaz de repositorio
    - Agrega consulta de repositorio simulado
    - Agrega seed de datos
20. Agrega Infraestructura Web (API REST)

### Otros temas de dominio a abordar
- Comunicación entre casos de uso (Mediator)
- Contenedor de acceso a los repositorios (Repository Manager)
- Agrupar transacciones de escritura en una única transacción (Unit of Work)
- Eventos de Dominio (Domain Events)

### Otros temas de API a abordar
- Versionamiento de API
- Paginación y ordenación
- Logging
- Error Handling
- Middlewares
- Hooks
- Cache
    - In-Memory
    - Distributed
- Tracing
- Timeout y Retries
- Health Check
- Rate Limit
- Circuit Breaker
- Pruebas de integración
- Pruebas de carga
- Documentación de API


## Postman (ES)
El proyecto 12 tiene infraestructura web (API REST) y se puede probar con Postman. La colección a importar está en la carpeta `postman`.

```bash
curl -X GET "http://localhost:8080/api/v1/products?min_price=0&max_price=200" -H "accept: application/json"
```