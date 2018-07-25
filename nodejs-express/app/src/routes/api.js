const express = require('express')

const router = express.Router()

router.get('/', (req, res) => res.send({
    message: 'Hello world from Hasura'
}))

module.exports = {router};
