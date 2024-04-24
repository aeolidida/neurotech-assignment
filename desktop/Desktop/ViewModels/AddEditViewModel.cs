using ReactiveUI;
using System;
using System.Collections.ObjectModel;
using System.Windows.Input;
using Desktop.Models;
using System.Reactive;
using System.Linq;

namespace Desktop.ViewModels
{
    public class AddEditViewModel : ReactiveObject
    {
        private Patient? _patient;
        private string? _fullName;
        private DateTimeOffset _birthday;

        private string? _selectedGender;
        private ObservableCollection<string>? _genders;

        public Patient? Patient
        {
            get => _patient;
            set => this.RaiseAndSetIfChanged(ref _patient, value);
        }

        public string? FullName
        {
            get => _fullName;
            set => this.RaiseAndSetIfChanged(ref _fullName, value);
        }

        public DateTimeOffset Birthday
        {
            get => _birthday;
            set => this.RaiseAndSetIfChanged(ref _birthday, value);
        }

        public string? SelectedGender
        {
            get => _selectedGender;
            set => this.RaiseAndSetIfChanged(ref _selectedGender, value);
        }

        public ObservableCollection<string>? Genders
        {
            get => _genders;
            set => this.RaiseAndSetIfChanged(ref _genders, value);
        }

        public ReactiveCommand<Unit, bool> SaveCommand { get; }

        public AddEditViewModel(Patient? patient)
        {
            Patient = patient?.Clone() ?? new Patient();

            FullName = Patient?.FullName;
            Birthday = new DateTimeOffset(Patient?.Birthday ?? DateTime.MinValue, TimeSpan.Zero);

            SelectedGender = Patient != null ? Patient.Gender.ToString() : Gender.Unknown.ToString();
            Genders = new ObservableCollection<string>(Enum.GetNames(typeof(Gender)).Where(g => g != Gender.Unknown.ToString())); 

            SaveCommand = ReactiveCommand.Create<Unit, bool>(SavePatient);
        }

        private bool SavePatient(Unit unit)
        {
            if (Patient != null)
            {
                Patient.FullName = FullName ?? string.Empty;
                Patient.Gender = (Gender)Enum.Parse(typeof(Gender), SelectedGender ?? Gender.Unknown.ToString());
                Patient.Birthday = Birthday.DateTime;

                if (!string.IsNullOrWhiteSpace(FullName) && Patient.Birthday != DateTime.MinValue)
                {
                    return true;
                }
            }
            
            return false;
        }
    }
}