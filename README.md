# Send mail with attachment

This project is able to send an email with an **attached file**.

The **addr variable** it has to be filled with the **SMTP server** (with port). For example:

```
addr := "smtp.gmail.com:587"
```

If **auth is needed**, a **host variable** has to be created and filled with **addr variable without the port**. For example:

```
host := "smtp.gmail.com"
```

Then use this two lines to get authorization and send the email:

```
auth := smtp.PlainAuth("", user, password, host)
err := smtp.SendMail(addr, auth, from, to, data)
```

*If you want to try with the **SMTP of Google**, **2-Step authentication** is needed and a password for apps has to be created.
Auth is needed, so **user variable** is the email and **pass variable** is the pass that Google will give you.*

That's all, enjoy.
