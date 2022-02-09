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
    Button1: TButton;
    Button2: TButton;
    ListView1: TListView;
    LeftBox: TGroupBox;
    MainMenu: TMainMenu;
    MenuItem1: TMenuItem;
    Maddpic: TMenuItem;
    Maddvd: TMenuItem;
    MenuItem10: TMenuItem;
    MenuItem11: TMenuItem;
    PVF: TMenuItem;
    PVD: TMenuItem;
    PVU: TMenuItem;
    MenuItem2: TMenuItem;
    PPF: TMenuItem;
    PPD: TMenuItem;
    PPU: TMenuItem;
    MenuItem8: TMenuItem;
    MenuItem9: TMenuItem;
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

end.

