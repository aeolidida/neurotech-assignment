using System;
using System.Threading.Tasks;
using Avalonia.Controls;
using Avalonia.Markup.Xaml;
using Desktop.Models;
using Desktop.ViewModels;

namespace Desktop.Views
{
    public partial class MainWindow : Window
    {
        private readonly MainWindowViewModel _viewModel;
        public MainWindow(MainWindowViewModel viewModel)
        {
            InitializeComponent();
            _viewModel = viewModel;
            this.DataContext = _viewModel;
        }

        private void InitializeComponent()
        {
            AvaloniaXamlLoader.Load(this);
        }

        private async void DeleteButton_Click(object sender, Avalonia.Interactivity.RoutedEventArgs e)
        {
            var button = (Button)sender;
            if (button.CommandParameter is not Guid guid) return;

            var confirmationWindow = new ConfirmationWindow
            {
                DataContext = new ConfirmationWindowViewModel(this),
            };
            var result = await confirmationWindow.ShowDialog<bool?>(this);

            if (result == true)
            {
                _viewModel.DeleteCommand.Execute(guid);
            }
        }

        private async void EditButton_Click(object sender, Avalonia.Interactivity.RoutedEventArgs e)
        {
            var button = (Button)sender;
            if (button.CommandParameter is not Patient patient) return;

            var addEditPatientWindow = new AddEditPatientWindow
            {
                DataContext = new AddEditViewModel(patient.Clone()) 
            };

            var result = await addEditPatientWindow.ShowDialog<bool?>(this);

            if (result == true)
            {
                var editedPatient = (addEditPatientWindow.DataContext as AddEditViewModel)?.Patient;
                if (editedPatient != null)
                {
                    _viewModel.EditCommand.Execute(editedPatient);
                }
            }
        }

        private async void AddButton_Click(object sender, Avalonia.Interactivity.RoutedEventArgs e)
        {
            var addEditPatientWindow = new AddEditPatientWindow
            {
                DataContext = new AddEditViewModel(null) 
            };

            var result = await addEditPatientWindow.ShowDialog<bool?>(this);

            if (result == true)
            {
                var addedPatient = (addEditPatientWindow.DataContext as AddEditViewModel)?.Patient;
                if (addedPatient != null)
                {
                    _viewModel.AddCommand.Execute(addedPatient);
                }
            }
        }
    }
}