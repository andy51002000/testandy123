<template>
  <div>
    <h2>Member Details</h2>
    <div v-if="loading">Loading...</div>
    <div v-else>
      <p><strong>Member ID:</strong> {{ member.pk }}</p>
      <p><strong>Username:</strong> {{ member.username }}</p>
      <p><strong>Creation Time:</strong> {{ member.create_time }}</p>
      <p><strong>Total Fee Last Year:</strong> {{ member.total_fee_last_year }}</p>
    </div>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  data() {
    return {
      loading: true,
      member: null,
    };
  },
  mounted() {
    // Extract member ID from route params
    const memberId = this.$route.params.id;
    
    // Make API request to fetch member details
    axios.get(`http://goserver:3000/member/${memberId}`)
      .then(response => {
        this.member = response.data;
        this.loading = false;
      })
      .catch(error => {
        console.error('Error fetching member details:', error);
        this.loading = false;
      });
  }
};
</script>

<style scoped>
/* Add styling as needed */
</style>
