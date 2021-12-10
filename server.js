const express = require("express")
const app = express()

app.get('/', (request, response) => {
    response.send("Submariners Server")
})

app.listen(3000, () => {
    console.log("Server started on port 3000.")
})
