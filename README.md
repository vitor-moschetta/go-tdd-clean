## Trilha

01. Input basico e UseCase basico retornando true
02. Input com Validação
03. Agrega Entidade e Interface do Repositorio. 
    - Acoplamento, injeção de dependencia e inversão de controle. 
    - Aqui vai dar erro de referência porque não temos uma implementação do repositório.
04. Agregamos um Repositorio em Memoria (mock)
    - Vamos ter erro de referência circular entre o arquivo de teste e o arquivo de mock. Necessário criar pacote para arquivo de teste.
05. Agregamos Validação de Entidade. Aqui podemos falar sobre validação de entidade e validação de input
    - Porque não validamos somente a entidade?         
        - Porque precisamos fornecer feedback para o usuário sobre os seus dados de entrada (Isso dará abertura para entrarmos com o conceito de Notification Pattern)
        - Porque a validação de entidade pode envolver regras mais complexas que a validação de input.
        - Fail Fast
06. Agregamos propriedade Price na Entidade
    - Fazemos uma introdução sobre Notification Pattern: Na primeira requisição o usuário precisa saber todas as falhas que ocorreram.
07. Agregamos Notification Pattern com Output do tipo []errors
08. Agregamos Notification Pattern com Output do tipo struct