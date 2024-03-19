<template>
  <div>
    <h1>Member Transaction</h1>
    <!-- Input fields to enter the member ID, start time, and end time -->
    <label for="memberIdInput">Enter Member ID:</label>
    <input type="text" id="memberIdInput" v-model="memberId">
    <label for="startTimeInput">Enter Start Time:</label>
    <input type="text" id="startTimeInput" v-model="startTime">
    <label for="endTimeInput">Enter End Time:</label>
    <input type="text" id="endTimeInput" v-model="endTime">
    <!-- Button to trigger fetching transactions -->
    <button @click="fetchTransactions">Fetch Transactions</button>
    <ul>
      <li v-for="transaction in transactions" :key="transaction.pk">
        <p>Transaction ID: {{ transaction.pk }}</p>
        <p>Member ID: {{ transaction.member_fk }}</p>
        <p>Type: {{ transaction.type }}</p>
        <p>Borrow Fee: {{ transaction.borrow_fee }}</p>
        <p>Create Time: {{ transaction.create_time }}</p>
      </li>
    </ul>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  name: 'TransactionPage',
  data() {
    return {
      memberId: '', // Member ID obtained from the input field
      startTime: '', // Start time obtained from the input field
      endTime: '', // End time obtained from the input field
      transactions: [],
    };
  },
  methods: {
    fetchTransactions() {
      // Construct URL with parameters
      const url = `https://edb0-104-198-4-130.ngrok-free.app/transactions/${this.memberId}?startTime=${this.startTime}&endTime=${this.endTime}`;

      // Make GET request
      axios.get(url)
        .then(response => {
          console.log('Response:', response.data);
          // Update transactions data with response
          this.transactions = response.data;
        })
        .catch(error => {
          console.error('Error fetching data:', error);
          console.error('Error details:', error.response); // Log the detailed error response
        });
    },
  },
};
</script>

<style scoped>
/* Component styles */
</style>
