<script>
    import { onMount } from "svelte";
    import axios from "$lib/axios"
    export let data;
    let res, userId, ticketId, failed, success;
onMount(async () => {
    try{
        const response = await axios.get(`/bought/ticket/${data.data}`);
        res = response.data;
        ticketId = res.TicketId;
        userId = res.UserId;
        console.log(res);
    }catch (e){
        console.error("Error:", e);
        failed = e.response.data.error;
    }
})

function updateBoughtTicketData(){
    if(userId == "" || userId == undefined || userId < 1){
        failed = "Users Id cant be empty or less than 1;";
    }else{
        failed = "";
    }
    if(ticketId == "" || ticketId == undefined || ticketId < 1){
        failed = "Ticket Id cant be empty or less than 1;";
    }else{
        failed = "";
    }
    axios.put(`/bought/ticket/${data.data}`, {
           ticketId: ticketId,
           userId: userId,
        })
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
    {#if res != "" && res != undefined}
        <div class="margin-t flex col">
            <label for="">
                Users Id:
                <input type="number" min="1" step="1" bind:value={userId}>
            </label>
            <label for="">
                Tickets id:
                <input type="number" min="1" step="1" bind:value={ticketId}>
            </label>
            {#if success != "" && success != undefined}
                <p class="success">{success}</p>
            {/if}
            {#if failed != "" && failed != undefined}
                <p class="error">{failed}</p>
            {/if}
            <button on:click={updateBoughtTicketData}>Uptete Bought Ticket Data</button>
            <!-- <img src={"http://localhost:8000/"+res.QrPath} alt="">
            <p>{res.UserName}</p>
            <p>{res.EventName}</p>
            <p>{res.TicketDate}</p>
            <p>{res.TicketLocation}</p>
            <p>{res.TicketSeat}</p> -->
        </div>
    {/if}
</main>

<style>
    .margin-t{
        margin-top: 5rem;
    }
    input {
        padding: 1rem;
        border: none;
        outline: none;
        border-radius: 0.2rem;
        background-color: var(--bg);
        font-size: 12pt;
        margin: 1rem;
    }
    button {
        background-color: var(--bg);
        border-radius: 0.4rem;
        padding: 0.7rem 5.7rem 0.7rem 5.7rem;
        border: none;
        outline: none;
    }
</style>