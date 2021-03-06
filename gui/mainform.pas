unit Mainform;

{$mode objfpc}{$H+}

interface

uses
  Classes, SysUtils, Forms, Controls, Graphics, Dialogs, StdCtrls, ComCtrls,
  ExtCtrls, Buttons, Menus;

type

  { TMainform }

  TMainform = class(TForm)
    Background: TImage;
    ComboBox1: TComboBox;
    ComboBox2: TComboBox;
    GroupBox1: TGroupBox;
    RadioButton1: TRadioButton;
    RBSelf: TRadioButton;
    StaticText1: TStaticText;
    StaticText2: TStaticText;
    SysGroup: TGroupBox;
    LogList: TListBox;
    StartBtn: TImage;
    Sysmsg: TStaticText;
    TaskList: TListView;
    LeftBox: TGroupBox;
    MainMenu: TMainMenu;
    MenuItem1: TMenuItem;
    Maddpic: TMenuItem;
    Maddvd: TMenuItem;
    PDel: TMenuItem;
    PDelAll: TMenuItem;
    PVF: TMenuItem;
    PVD: TMenuItem;
    PVU: TMenuItem;
    MenuItem2: TMenuItem;
    PPF: TMenuItem;
    PPD: TMenuItem;
    PPU: TMenuItem;
    MenuItem8: TMenuItem;
    PReload: TMenuItem;
    MVU: TMenuItem;
    MPF: TMenuItem;
    MenuItem4: TMenuItem;
    MenuItem5: TMenuItem;
    MPD: TMenuItem;
    MPU: TMenuItem;
    MVF: TMenuItem;
    MVD: TMenuItem;
    PopupList: TPopupMenu;
    RightBox: TGroupBox;
    PBpres: TProgressBar;
    Pres: TStaticText;
    procedure BackgroundClick(Sender: TObject);
    procedure FormClose(Sender: TObject; var CloseAction: TCloseAction);
    procedure FormCreate(Sender: TObject);
    procedure FormShow(Sender: TObject);
    procedure LogListClick(Sender: TObject);
    procedure PDelAllClick(Sender: TObject);
    procedure PDelClick(Sender: TObject);
    procedure PopupListPopup(Sender: TObject);
    procedure StartBtnMouseDown(Sender: TObject; Button: TMouseButton;
      Shift: TShiftState; X, Y: integer);
    procedure StartBtnMouseLeave(Sender: TObject);
    procedure StartBtnMouseMove(Sender: TObject; Shift: TShiftState; X, Y: integer);
    procedure StartBtnMouseUp(Sender: TObject; Button: TMouseButton;
      Shift: TShiftState; X, Y: integer);
    procedure SysmsgClick(Sender: TObject);
    procedure TaskListClick(Sender: TObject);
    procedure TaskListMouseDown(Sender: TObject; Button: TMouseButton;
      Shift: TShiftState; X, Y: integer);
    procedure TaskListMouseLeave(Sender: TObject);
    procedure TaskListMouseMove(Sender: TObject; Shift: TShiftState; X, Y: integer);
    procedure MenuItem1Click(Sender: TObject);
    procedure MaddpicClick(Sender: TObject);
    procedure MaddvdClick(Sender: TObject);
    procedure PPFClick(Sender: TObject);
    procedure MPDClick(Sender: TObject);
    procedure MPFClick(Sender: TObject);
    procedure MenuItem5Click(Sender: TObject);
    procedure MPUClick(Sender: TObject);
    procedure MVDClick(Sender: TObject);
    procedure MVFClick(Sender: TObject);
    procedure MVUClick(Sender: TObject);
    procedure TaskListMouseUp(Sender: TObject; Button: TMouseButton;
      Shift: TShiftState; X, Y: integer);
  private

  public

  end;

var
  Mainform: TMainform;

implementation

{$R *.lfm}

{ TMainform }

procedure TMainform.FormCreate(Sender: TObject);
begin

end;

procedure TMainform.BackgroundClick(Sender: TObject);
begin

end;

procedure TMainform.FormClose(Sender: TObject; var CloseAction: TCloseAction);
begin

end;

procedure TMainform.FormShow(Sender: TObject);
begin

end;

procedure TMainform.LogListClick(Sender: TObject);
begin

end;

procedure TMainform.PDelAllClick(Sender: TObject);
begin

end;

procedure TMainform.PDelClick(Sender: TObject);
begin

end;

procedure TMainform.PopupListPopup(Sender: TObject);
begin

end;

procedure TMainform.StartBtnMouseDown(Sender: TObject; Button: TMouseButton;
  Shift: TShiftState; X, Y: integer);
begin

end;

procedure TMainform.StartBtnMouseLeave(Sender: TObject);
begin

end;

procedure TMainform.StartBtnMouseMove(Sender: TObject; Shift: TShiftState;
  X, Y: integer);
begin

end;

procedure TMainform.StartBtnMouseUp(Sender: TObject; Button: TMouseButton;
  Shift: TShiftState; X, Y: integer);
begin

end;

procedure TMainform.SysmsgClick(Sender: TObject);
begin

end;

procedure TMainform.TaskListClick(Sender: TObject);
begin

end;

procedure TMainform.TaskListMouseDown(Sender: TObject; Button: TMouseButton;
  Shift: TShiftState; X, Y: integer);
begin

end;

procedure TMainform.TaskListMouseLeave(Sender: TObject);
begin

end;

procedure TMainform.TaskListMouseMove(Sender: TObject; Shift: TShiftState;
  X, Y: integer);
begin

end;

procedure TMainform.MenuItem1Click(Sender: TObject);
begin

end;

procedure TMainform.MaddpicClick(Sender: TObject);
begin

end;

procedure TMainform.MaddvdClick(Sender: TObject);
begin

end;

procedure TMainform.PPFClick(Sender: TObject);
begin

end;

procedure TMainform.MPDClick(Sender: TObject);
begin

end;

procedure TMainform.MPFClick(Sender: TObject);
begin

end;

procedure TMainform.MenuItem5Click(Sender: TObject);
begin

end;

procedure TMainform.MPUClick(Sender: TObject);
begin

end;

procedure TMainform.MVDClick(Sender: TObject);
begin

end;

procedure TMainform.MVFClick(Sender: TObject);
begin

end;

procedure TMainform.MVUClick(Sender: TObject);
begin

end;

procedure TMainform.TaskListMouseUp(Sender: TObject; Button: TMouseButton;
  Shift: TShiftState; X, Y: integer);
begin

end;

end.



