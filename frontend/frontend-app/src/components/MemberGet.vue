<template>
  <div>
    <h2>Member Details</h2>
    <div>
      <label for="memberId">Enter Member ID:</label>
      <input type="text" id="memberId" v-model="memberId">
      <button @click="fetchMemberDetails">Fetch Details</button>
    </div>
    <div v-if="loading">Loading...</div>
    <div v-else-if="member">
      <p><strong>Member ID:</strong> {{ member.pk }}</p>
      <p><strong>Username:</strong> {{ member.username }}</p>
      <p><strong>Creation Time:</strong> {{ member.create_time }}</p>
      <p><strong>Total Fee Last Year:</strong> {{ member.total_fee_last_year }}</p>
    </div>
    <div v-else>
      <p>No member details available</p>
    </div>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  data() {
    return {
      loading: false,
      memberId: '', // Store member ID entered by the user
      member: null,
    };
  },
  methods: {
    fetchMemberDetails() {
      // Ensure member ID is not empty
      if (!this.memberId.trim()) {
        alert('Please enter a valid member ID');
        return;
      }
      
      this.loading = true;
      
      // Make API request to fetch member details
      axios.get(`https://edb0-104-198-4-130.ngrok-free.app/member/${this.memberId}`)
        .then(response => {
          this.member = response.data;
          this.loading = false;
        })
        .catch(error => {
          console.error('Error fetching member details:', error);
          this.loading = false;
        });
    }
  }
};
</script>

<style scoped>
/* Add styling as needed */
</style>
