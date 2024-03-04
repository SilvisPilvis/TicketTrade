<script>
    import { onMount } from "svelte";
    import axios from "$lib/axios"
    export let data;
    let res;
onMount(async () => {
    try{
        const response = await axios.get(`/${data.data}/tickets`);
        res = response.data;
        console.log(res);
    }catch (e){
        console.error("Error:", e);
        failed = e.response.data.error;
    }
})
</script>

<main class="flex col cen">
    {#if res != "" && res != undefined}
    <h1 class="margin-t">Your Tickets:</h1>
        <div class="flex cen row">
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
        </div>
    {/if}
</main>

<style>
    .margin-t{
        margin-top: 5rem;
    }
    .row{
        flex-direction: row;
        flex-wrap: wrap;
    }
    .ticket{
        background-color: var(--fg);
        width: 15rem;
        border-radius: 0.4rem;
        margin: 1rem;
        padding: 1rem;
    }
    .name{
        font-weight: 900;
    }
</style>