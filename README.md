# vpn-mfa.exe

Esse script tem o propósito de realizar reconexão automática com o openvpn
toda vez que houver desconexão por conta do MFA.

## Passos

- Adicionar comando 'openvpn' no terminal
  -Adicionar caminho C:\Program Files\OpenVPN\bin a variável PATH.

- Obter secret de QRCode e criar variável de ambiente VPN_MFA

  - https://scanqr.org/
  - No site fazer upload da imagem do seu QRCode da VPN.
  - Obter código que aparece na tag 'secret'.
  - Criar variável de ambiente VPN_MFA com o código obtido.

- Criar variável de ambiente VPN_PREFIX

  - Valor deve ser o prefixo da senha que você utiliza.

- Criar variável de ambiente VPN_USER

  - Valor deve ser seu usuário do VPN.
  - Provavelmente nome.sobrenome

- Criar variável de ambiente VPN_CONFIG_PATH

  - Inserir caminho do arquivo do openvpn.

- Baixar ou gerar executável vpn-mfa.exe e torná-lo disponível como comando
  Para conectar na VPN com um simples comando no terminal
  deve-se baixar o executável, colocar em uma pasta de sua escolha
  e concatenar o endereço da pasta na variável de ambiente PATH.

  - Isso vai tornar o comando 'vpn-mfa' disponìvel no terminal.

- Abrir terminal em modo administrador
- Executar 'vpn-mfa'
  - O comando irá reconectar automaticamente toda vez que a VPN cair.
