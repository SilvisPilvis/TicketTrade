<script>
	// import Stripe from "stripe";
	// export const stripe = new Stripe(import.meta.env.VITE_STRIPE_API_KEY);
	import { onMount } from "svelte";
	import {loadStripe} from '@stripe/stripe-js';
	import Cookies from "js-cookie";
	import axios from "axios";
	export let data;
	let stripe, success;
	onMount(async () => {
		// const stripe = await loadStripe(import.meta.env.VITE_STRIPE_API_KEY);
		stripe = await loadStripe(import.meta.env.VITE_STRIPE_API_KEY);
		checkout();
	})
	
	
	// let cardNumber, expiry, cvc, email;
	function checkout(){
		if(data.data == "" || data.data == undefined){
			console.log("Id is undefined.");
			return;
		}
		// let value = {
		// 	email: email
		// }
		// console.log(value);
		const config = {
            headers: { Authorization: `Bearer ${Cookies.get('token')}` }
        };
		axios.post(`http://localhost:8000/checkout/${data.data}`, {}, config)
        .then(function (response) {
            console.log(response);
            success = response.data;
			console.log(success);
			const id = success.Id;
			stripe.redirectToCheckout({sessionId: id});
        })
        .catch(function (error) {
            console.log(error);
            failed = error.response.data.error;
        });
		// fetch("http://localhost:8000/checkout", {
		// 	method: "POST",
		// }).then(
		// 	async response => {
		// 		const res = await response.json();
		// 		console.log(res);
		// 		const id = res.Id;
		// 		stripe.redirectToCheckout({sessionId: id});
		// 	}
		// )
	}
</script>