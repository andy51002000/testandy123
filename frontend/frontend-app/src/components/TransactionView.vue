<template>
  <div>
    <h1>Member Transaction</h1>
    <!-- Your component content here -->
  </div>
</template>

<script>
import axios from 'axios';

export default {
  name: 'TransactionPage',
  data() {
    return {
      transactions: [],
    };
  },
  mounted() {
    this.fetchTransactions();
  },
  methods: {
    fetchTransactions() {
      // Extract parameters from the URL using $route.params
      const memberId = this.$route.params.memberId;
      const startTime = this.$route.query.startTime;
      const endTime = this.$route.query.endTime;

      // Construct URL with parameters
      const url = `http://goserver:3000/transactions/${memberId}?startTime=${startTime}&endTime=${endTime}`;

      // Make GET request
      axios.get(url)
        .then(response => {
          this.transactions = response.data;
        })
        .catch(error => {
          console.error('Error fetching transactions:', error);
        });
    },
  },
};
</script>

<style scoped>
/* Component styles */
</style>