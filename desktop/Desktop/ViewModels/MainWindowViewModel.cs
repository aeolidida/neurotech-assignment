using System;
using System.Collections.ObjectModel;
using System.Linq;
using System.Net.Http;
using System.Reactive.Linq;
using System.Threading.Tasks;
using System.Windows.Input;
using Avalonia.Controls;
using Desktop.Models;
using Desktop.Services;
using Desktop.Views;
using ReactiveUI;

namespace Desktop.ViewModels
{
    public class MainWindowViewModel : ReactiveObject
    {
        private readonly ApiService _apiService;
        private ObservableCollection<Patient> _patients = new ObservableCollection<Patient>();
        private ObservableCollection<Patient> _filteredPatients = new ObservableCollection<Patient>();
        private string _statusMessage = string.Empty;
        private string _searchText = string.Empty;

        public ObservableCollection<Patient> Patients
        {
            get => _patients;
            set => this.RaiseAndSetIfChanged(ref _patients, value);
        }

        public ObservableCollection<Patient> FilteredPatients
        {
            get => _filteredPatients;
            set => this.RaiseAndSetIfChanged(ref _filteredPatients, value);
        }

        public string SearchText
        {
            get => _searchText;
            set
            {
                this.RaiseAndSetIfChanged(ref _searchText, value);
                FilterPatients();
            }
        }

        public string StatusMessage
        {
            get => _statusMessage;
            set => this.RaiseAndSetIfChanged(ref _statusMessage, value);
        }

        public ICommand AddCommand { get; }
        public ICommand EditCommand { get; }
        public ICommand DeleteCommand { get; }
        public ICommand RefreshCommand { get; }

        public MainWindowViewModel()
        {
            _apiService = new ApiService();
            Patients = new ObservableCollection<Patient>();

            AddCommand = ReactiveCommand.CreateFromTask<Patient>(AddPatient);
            EditCommand = ReactiveCommand.CreateFromTask<Patient>(EditPatient);
            DeleteCommand = ReactiveCommand.CreateFromTask<Guid>(DeletePatient);
            RefreshCommand = ReactiveCommand.CreateFromTask(RefreshPatients);

            _ = RefreshPatients();
        }

        private async Task AddPatient(Patient patient)
        {
            try
            {
                var newPatient = await _apiService.NewPatient(patient);
                Patients.Add(newPatient);
                FilterPatients();
            }
            catch (HttpRequestException ex)
            {
                StatusMessage = $"Error adding patient: {ex.Message}";
            }
        }

        private async Task EditPatient(Patient patient)
        {
            try
            {
                var editedPatient = await _apiService.EditPatient(patient);

                var foundPatient = Patients.FirstOrDefault(p => p.GUID == editedPatient.GUID);
                if (foundPatient != null)
                {
                    var index = Patients.IndexOf(foundPatient);
                    if (index != -1)
                    {
                        Patients[index] = editedPatient;
                        FilterPatients();
                    }
                }
                
            }
            catch (HttpRequestException ex)
            {
                StatusMessage = $"Error editing patient: {ex.Message}";
            }
        }

        private async Task DeletePatient(Guid guid)
        {
            try
            {
                var selectedPatient = Patients.FirstOrDefault(p => p.GUID == guid);
                if (selectedPatient is null) return;

                await _apiService.DeletePatient(selectedPatient.GUID.ToString());
                Patients.Remove(selectedPatient);
                FilterPatients();
            }
            catch (HttpRequestException ex)
            {
                StatusMessage = $"Error deleting patient: {ex.Message}";
            }
        }

        private async Task RefreshPatients()
        {
            try
            {
                StatusMessage = "Loading patients...";
                var patients = await _apiService.GetListPatients();
                Patients = new ObservableCollection<Patient>(patients);
                FilterPatients();
                StatusMessage = string.Empty;
            }
            catch (HttpRequestException ex)
            {
                StatusMessage = $"Error loading patients: {ex.Message}";
            }
        }
        private void FilterPatients()
        {
            if (string.IsNullOrWhiteSpace(SearchText))
            {
                FilteredPatients = new ObservableCollection<Patient>(Patients);
                return;
            }

            FilteredPatients = new ObservableCollection<Patient>(
                Patients.Where(p => p.FullName.ToLower().Contains(SearchText.ToLower()))
            );
        }
    }
}