<template>
  <div>
    <h2>Update Member</h2>
    <form @submit.prevent="updateMember">
      <label for="memberId">Member ID:</label>
      <input type="text" id="memberId" v-model="memberId" required>
      <label for="newUsername">New Username:</label>
      <input type="text" id="newUsername" v-model="newUsername" required>
      <button type="submit">Update Member</button>
    </form>
    <p v-if="message">{{ message }}</p>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  data() {
    return {
      memberId: '',
      newUsername: '',
      message: ''
    };
  },
  methods: {
    updateMember() {
      // Prepare data to be sent in the request
      const requestData = {
        username: this.newUsername
      };

      // Make API request to update the member
      axios.put(`https://edb0-104-198-4-130.ngrok-free.app/member/${this.memberId}`, requestData)
        .then(response => {
          this.message = response.data.message;
        })
        .catch(error => {
          console.error('Error updating member:', error);
          this.message = 'Failed to update member';
        });
    }
  }
};
</script>

<style scoped>
/* Add styling as needed */
</style>
