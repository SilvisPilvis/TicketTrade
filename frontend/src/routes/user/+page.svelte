<script>
import axios from "$lib/axios";
import Cookies from "js-cookie";
// import {onMount} from "svelte";
// export let data;
let changedPass, changedUsername, failed, success;

function changeUsername(){
    const config = {
        headers: { Authorization: `Bearer ${Cookies.get('token')}` }
    };
    axios.put('/username', {
            Username: changedUsername
        }, config)
        .then(function (response) {
            console.log(response);
            success = response.data
        })
        .catch(function (error) {
            console.log(error);
            failed = error.response.data.error;
        });
}
function changePassword(){
    const config = {
        headers: { Authorization: `Bearer ${Cookies.get('token')}` }
    };
    axios.put('/password', {
            Username: "a",
            Password: changedPass
        }, config)
        .then(function (response) {
            console.log(response);
            success = response.data
        })
        .catch(function (error) {
            console.log(error);
            failed = error.response.data.error;
        });
}
</script>

<main class="flex col cen">
    <div class="margin-t flex col">
        <label class="flex col">
            Change Username:
            <input type="text" bind:value={changedUsername}/>
            <button on:click={changeUsername}>Change Username</button>
        </label>
        <label class="flex col">
            Change password:
            <input type="password" bind:value={changedPass}/>
            <button on:click={changePassword}>Change Password</button>
        </label>
        {#if failed != "" && failed != undefined}
            <p class="error">{failed}</p>
        {/if}
        {#if success != "" && success != undefined}
            <p class="error">{success}</p>
        {/if}
    </div>
</main>

<style>
    .margin-t{
        margin-top: 5rem;
    }
    label{
        margin: 1rem;
    }
    button, input{
        display: flex;
        margin: 1rem;
        padding: 0.7rem;
        border-radius: 0.4rem;
        background-color: var(--bg);
        border: none;
    }
</style>