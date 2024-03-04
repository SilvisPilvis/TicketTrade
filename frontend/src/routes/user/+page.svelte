<script>
import axios from "$lib/axios";
import Cookies from "js-cookie";
// import {onMount} from "svelte";
// export let data;
let changedPass, changedUsername, failed, success;

// onMount(async () => {
//     try{
//         const response = await axios.get(`/${data.data}/tickets`);
//         res = response.data;
//         console.log(res);
//     }catch (e){
//         console.error("Error:", e);
//         failed = e.response.data.error;
//     }
// });

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
    <!-- <div class="flex cen row">
        {#each res as ticket}
            <div class="ticket">
                <p class="name">{ticket.EventName}</p>
                <p>User: {ticket.UserName}</p>
                <p>{ticket.TicketBoughtAt}</p>
                <p>{ticket.TicketLocation}</p>
                <p>{ticket.TicketDate}</p>
                <p>{ticket.TicketSeat}</p>
            </div>
        {/each}
    </div> -->
    <div class="margin-t flex col cen">
        <label class="flex col cen">
            Change Username:
            <input type="text" bind:value={changedUsername}/>
            <button on:click={changeUsername}>Change Username</button>
        </label>
        <label class="flex col cen">
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
    main{
        height: 100%;
    }
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
        width: 16rem;
        border-radius: 0.4rem;
        background-color: var(--fg);
        border: none;
    }
    button{
        display: flex;
        align-items: center;
        justify-content: center;
        /* text-align: center; */
        background-color: var(--button-fill);
    }
</style>