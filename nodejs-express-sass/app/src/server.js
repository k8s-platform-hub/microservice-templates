const express = require('express')
const path = require('path')

const app = express()

// Redirect all static files to public/
app.use('/static', express.static(path.join(__dirname, 'public')))

const homePath = path.join(__dirname, 'public', 'html', 'index.html')
const notFoundPath = path.join(__dirname, 'public', 'html', '404.html')
const port = process.env.APP_PORT || '8080'

app.get('/', (req, res) => res.sendFile(homePath))
app.all('*', (req, res) => res.sendFile(notFoundPath))

app.listen(port, () => console.log(`Example app listening on port ${port}!`))
