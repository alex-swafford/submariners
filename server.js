const express = require('express')
const path = require('path')

const app = express()

app.use(express.static(path.resolve(__dirname, 'client/build')))

app.get('*', (request, response) => {
    response.sendFile(path.resolve(__dirname, 'client/build', 'index.html'))
})

app.listen(3000, () => {
    console.log("Server started on port 3000.")
})
