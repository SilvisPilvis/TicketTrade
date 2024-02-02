<script>
import jwt from 'jsonwebtoken';
let token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFeHBpcmVzQXQiOjE3MDY4NTg1NDAsIklhdCI6MTcwNjg1ODI0MCwicm9sZSI6MSwidXNlcklkIjoxLCJ1c2VybmFtZSI6InRlc3QifQ.8Ict4OJobBeivkOGSLAW3EndX5BQHyAHVSlj7Xb8csg";
let secret = "S39WQoFWOCTw+obda12jg3dZ5dt7OZmpKdLmSZ83GiiVXVT/7NQIpq3XvgAgL4HV";
let decoded = jwt.verify(token, secret);
// console.log(decoded);
</script>

<!-- <main>
    <a href="/"></a>
</main> -->