using System.Reactive;
using System.Reactive.Linq;
using Avalonia;
using Avalonia.Controls;
using Avalonia.Markup.Xaml;
using Desktop.ViewModels;

namespace Desktop.Views
    {
    public partial class AddEditPatientWindow : Window
    {
        public AddEditPatientWindow()
        {
            InitializeComponent();
    #if DEBUG
            this.AttachDevTools();
    #endif
        }

        private void InitializeComponent()
        {
            AvaloniaXamlLoader.Load(this);
        }

        private async void OnSaveClick(object sender, Avalonia.Interactivity.RoutedEventArgs e)
        {
            var viewModel = DataContext as AddEditViewModel;
            if (viewModel != null)
            {
                var result = await viewModel.SaveCommand.Execute(Unit.Default);
                if (result)
                {
                    Close(true);
                }
                else
                {
                    var errorWindow = new ErrorWindow("Пожалуйста, заполните все поля.");
                    await errorWindow.ShowDialog(this);
                }
            }
        }

        private void OnCancelClick(object sender, Avalonia.Interactivity.RoutedEventArgs e)
        {
            Close(false);
        }
        
    }
}