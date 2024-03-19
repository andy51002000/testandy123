<template>
  <div>
    <h2>Update Member</h2>
    <form @submit.prevent="updateMember">
      <label for="username">New Username:</label>
      <input type="text" id="username" v-model="newUsername" required>
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
      newUsername: '',
      message: ''
    };
  },
  methods: {
    updateMember() {
      const memberId = this.$route.params.id; // Extract member ID from route params

      // Prepare data to be sent in the request
      const requestData = {
        username: this.newUsername
      };

      // Make API request to update the member
      axios.put(`http://goserver:3000/member/${memberId}`, requestData)
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
