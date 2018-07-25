const express = require('express')

const app = express()

app.use(require('./routes/web').router)
app.use('/api', require('./routes/api').router)

app.listen(8080, () => console.log('Example app listening on port 8080!'))
