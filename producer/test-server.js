// npm install express body-parser

const express = require('express');
const bodyParser = require('body-parser');
const app = express();
const port = 1234;

app.use(bodyParser.json());
app.use(function (req, res) {
    console.log(req.body);
    res.send('OK');
});

app.listen(port, () => {
    console.log(`Example app listening on port ${port}`);
});
