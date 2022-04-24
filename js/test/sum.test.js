const request = require('supertest');
const express = require('express');

const app = express();

app.get('/user', function (req, res) {
  res.status(200).json({ name: 'john' });
});




describe("test", () => {
  it("dsd", async () => {

    const res = await request(app)
      .get('/user')
      .send()

    expect(res.statusCode).toEqual(200)

  })
})