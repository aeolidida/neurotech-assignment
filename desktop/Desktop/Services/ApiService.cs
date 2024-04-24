using System;
using System.Collections.Generic;
using System.Collections.ObjectModel;
using System.Net.Http;
using System.Net.Http.Json;
using System.Text;
using System.Text.Json;
using System.Threading.Tasks;
using Desktop.Models;

namespace Desktop.Services
{
    public class ApiService
    {
        private readonly HttpClient _httpClient;
        private readonly string _baseUrl;

        public ApiService(string baseUrl = "http://localhost:8080")
        {
            _httpClient = new HttpClient();
            _baseUrl = baseUrl;
        }

        public async Task<IEnumerable<Patient>> GetListPatients()
        {
            try
            {
                var response = await _httpClient.GetAsync($"{_baseUrl}/getListPatients");
                response.EnsureSuccessStatusCode();

                var patients = await response.Content.ReadFromJsonAsync<IEnumerable<Patient>>();
                            
                if (patients != null)
                {
                    return patients;
                }
                {
                    return new ObservableCollection<Patient>();
                }
            }
            catch (HttpRequestException ex)
            {
                Console.WriteLine($"Error getting patients: {ex.Message}");
                throw;
            }
        }

        public async Task<Patient> NewPatient(Patient patient)
        {
            try
            {
                var json = JsonSerializer.Serialize(patient);
                var data = new StringContent(json, Encoding.UTF8, "application/json");

                var response = await _httpClient.PostAsync($"{_baseUrl}/newPatient", data);
                response.EnsureSuccessStatusCode();

                var newPatient = await response.Content.ReadFromJsonAsync<Patient>();
         
                if (newPatient != null)
                {
                    patient.GUID = newPatient.GUID;
                    return patient;
                }
                else
                {
                    throw new Exception($"Invalid or empty server response: {response.Content}");
                }

            }
            catch (HttpRequestException ex)
            {
                Console.WriteLine($"Error creating patient: {ex.Message}");
                throw;
            }
        }

        public async Task<Patient> EditPatient(Patient patient)
        {
            try
            {
                var json = JsonSerializer.Serialize(patient);
                var data = new StringContent(json, Encoding.UTF8, "application/json");

                var response = await _httpClient.PostAsync($"{_baseUrl}/editPatient", data);
                response.EnsureSuccessStatusCode();

                return patient;
            }
            catch (HttpRequestException ex)
            {
                Console.WriteLine($"Error editing patient: {ex.Message}");
                throw;
            }
        }

        public async Task DeletePatient(string guid)
        {
            try
            {
                var response = await _httpClient.PostAsync($"{_baseUrl}/delPatient?guid={guid}", null);
                response.EnsureSuccessStatusCode();
            }
            catch (HttpRequestException ex)
            {
                Console.WriteLine($"Error deleting patient: {ex.Message}");
                throw;
            }
        }
    }
}