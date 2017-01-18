#pragma once
#include <stdio.h>
namespace LoginPromt {
	using namespace System;
	using namespace System::ComponentModel;
	using namespace System::Collections;
	using namespace System::Windows::Forms;
	using namespace System::Data;
	using namespace System::Drawing;
	using namespace System::Diagnostics;
	using namespace System::IO;
	/// <summary>
	/// Summary for MyForm
	/// </summary>
	public ref class MyForm : public System::Windows::Forms::Form
	{
	public:
		MyForm(void)
		{
			InitializeComponent();
			//
			//TODO: Add the constructor code here
			//
		}

	protected:
		/// <summary>
		/// Clean up any resources being used.
		/// </summary>
		~MyForm()
		{
			if (components)
			{
				delete components;
			}
		}
	private: System::Windows::Forms::Label^  label1;
	protected:

	private: System::Windows::Forms::Button^  Exit;
	private: System::Windows::Forms::Timer^  timer1;

	private: System::ComponentModel::IContainer^  components;
	private: static System::Windows::Forms::Label^  CountDown;
			 static System::Int32 count = 200;
			 static String FormatTime(Int32 secs)
			 {
				 Int32 mins, hours;
				 mins = secs / 60;
				 hours = mins / 60;
				 h = String::String("hours: ");
				 String s =  h + Convert::ToString(hours);
				 return 
			 }

	static void TimerEventProcessor(Object^ obj, EventArgs^ args)
	{	
		CountDown->Text = FormatTime(count--);	//Convert::ToString(count--);
	}

	protected:

	private:
		/// <summary>
		/// Required designer variable.
		/// </summary>


#pragma region Windows Form Designer generated code
		/// <summary>
		/// Required method for Designer support - do not modify
		/// the contents of this method with the code editor.
		/// </summary>
		void InitializeComponent(void)
		{
			this->components = (gcnew System::ComponentModel::Container());
			this->label1 = (gcnew System::Windows::Forms::Label());
			this->Exit = (gcnew System::Windows::Forms::Button());
			this->timer1 = (gcnew System::Windows::Forms::Timer(this->components));
			this->CountDown = (gcnew System::Windows::Forms::Label());
			this->SuspendLayout();
			// 
			// label1
			// 
			this->label1->Dock = System::Windows::Forms::DockStyle::Top;
			this->label1->Font = (gcnew System::Drawing::Font(L"Times New Roman", 36, System::Drawing::FontStyle::Bold, System::Drawing::GraphicsUnit::Point,
				static_cast<System::Byte>(0)));
			this->label1->ForeColor = System::Drawing::Color::MediumBlue;
			this->label1->Location = System::Drawing::Point(0, 0);
			this->label1->Name = L"label1";
			this->label1->Size = System::Drawing::Size(1231, 71);
			this->label1->TabIndex = 0;
			this->label1->Text = L"ALL YOUR FILES ARE ENCRYPTED";
			this->label1->TextAlign = System::Drawing::ContentAlignment::MiddleCenter;
			this->label1->Click += gcnew System::EventHandler(this, &MyForm::label1_Click);
			// 
			// Exit
			// 
			this->Exit->AutoSize = true;
			this->Exit->Location = System::Drawing::Point(1088, 411);
			this->Exit->Name = L"Exit";
			this->Exit->Size = System::Drawing::Size(131, 33);
			this->Exit->TabIndex = 1;
			this->Exit->Text = L"Start Sesion";
			this->Exit->UseVisualStyleBackColor = true;
			this->Exit->Click += gcnew System::EventHandler(this, &MyForm::button1_Click);
			// 
			// timer1
			// 
			this->timer1->Enabled = true;
			this->timer1->Interval = 1000;
			this->timer1->Tick += gcnew System::EventHandler(this, &MyForm::TimerEventProcessor);
			// 
			// CountDown
			// 
			this->CountDown->Location = System::Drawing::Point(425, 71);
			this->CountDown->Name = L"CountDown";
			this->CountDown->Size = System::Drawing::Size(410, 62);
			this->CountDown->TabIndex = 2;
			this->CountDown->Text = L"CountDown";
			this->CountDown->Click += gcnew System::EventHandler(this, &MyForm::label2_Click);
			// 
			// MyForm
			// 
			this->AutoScaleDimensions = System::Drawing::SizeF(6, 13);
			this->AutoScaleMode = System::Windows::Forms::AutoScaleMode::Font;
			this->AutoSize = true;
			this->BackColor = System::Drawing::Color::LawnGreen;
			this->ClientSize = System::Drawing::Size(1231, 478);
			this->ControlBox = false;
			this->Controls->Add(this->CountDown);
			this->Controls->Add(this->Exit);
			this->Controls->Add(this->label1);
			this->FormBorderStyle = System::Windows::Forms::FormBorderStyle::None;
			this->Name = L"MyForm";
			this->ShowIcon = false;
			this->ShowInTaskbar = false;
			this->Text = L"Wellcome";
			this->WindowState = System::Windows::Forms::FormWindowState::Maximized;
			this->ResumeLayout(false);
			this->PerformLayout();

		}
#pragma endregion
	private: System::Void label1_Click(System::Object^  sender, System::EventArgs^  e) {
	}
	private: System::Void button1_Click(System::Object^  sender, System::EventArgs^  e) {

		//Process::Start("C:\\Windows\\System32\\cmd.exe");
		Application::Exit();
	}
	
	private: System::Void label2_Click(System::Object^  sender, System::EventArgs^  e) {
	}
};
}