package util

const EmailTemplate = `
<!DOCTYPE html>
<html lang="pt-br">
<head>
    <meta charset="UTF-8">
</head>
<html>
<body>

	<h2>Olá {{.Name}},</h2>
	<p>{{.Message}}</p>
	<p>Atenciosamente,<br>Equipe do Projeto Farmácia</p>
	<br />
	<img src="https://ik.imagekit.io/vzr6ryejm/email/footer_email.drawio.png?updatedAt=1726939931146" alt="footer" />
</body>
</html>
`