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

<main>
{#if result != "" && result != undefined}
    {#each result as ticket}
        <div>
            <p>{ticket.EventName}</p>
            <img src={"http://localhost:8000/"+ticket.TickedQRPath} alt="">
            <p>{ticket.TicketDate}</p>
            <p>{ticket.TicketLocation}</p>
            <p>{ticket.TicketSeat}</p>
        </div>
    {/each}
{/if}
</main>