<script>
import axios from '../services/axios';
let username, email, password, password_repeat, usernameError, emailError, passwordError, repeatError;

function register(){
    if(password !== password_repeat){
        repeatError = "Passwords must match";
        return;
    }
    axios.post('/register', {
        username: username,
        email: email,
        password: password
    })
    .then(function (response) {
        console.log(response.data.success);
    })
    .catch(function (error) {
        console.log(error);
    });
}
</script>

<main>
<input type="text" name="username" id="" bind:value={username} placeholder="username">
<input type="email" name="email" id="" bind:value={email} placeholder="email">
<input type="password" name="password" id="" bind:value={password} placeholder="password">
<input type="password" id="" bind:value={password_repeat} placeholder="password">
<button on:click={register}>Register</button>
</main>

<style>
:global(*){
    box-sizing: border-box;
    min-width: 0;
}
</style>