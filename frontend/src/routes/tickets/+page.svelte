<script>
    import axios from "$lib/axios";
    import { onMount } from "svelte";
    let result, failure; 
    onMount(async () => {
        try{
            const res = await axios.get('/tickets');
            // const res = await axios.get('/events');
            result = res.data;
            console.log(result)
        }catch (e){
            console.error("Error:", e.response.data.error);
            failure = e.response.data.error;
        }
    })
</script>

<main class="flex col cen">
{#if result != "" && result != undefined}
    <div class="margin-t flex cen row">
        {#each result as ticket}
            <div class="ticket">
                <p class="name">{ticket.EventName}</p>
                <img src={"http://localhost:8000/"+ticket.TickedQRPath} alt="">
                <p>{ticket.TicketDate}</p>
                <p>{ticket.TicketLocation}</p>
                <p>{ticket.TicketSeat}</p>
            </div>
        {/each}
    </div>
{/if}
</main>

<style>
    .row{
        flex-direction: row;
        flex-wrap: wrap;
    }
    .name{
        font-weight: 900;
    }
    .ticket{
        background-color: var(--bg);
        border-radius: 0.4rem;
        margin: 1rem;
        padding: 1rem;
    }
    .margin-t{
        margin-top: 5rem;
    }
</style>